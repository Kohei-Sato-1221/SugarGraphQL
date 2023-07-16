package graph

import (
	"context"
	"errors"

	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/graph/services"
	"github.com/graph-gophers/dataloader"
)

type Loaders struct {
	UserLoader *dataloader.Loader
}

func NewLoaders(Srv services.Services) *Loaders {
	userBatcher := &userBatcher{Srv: Srv}
	return &Loaders{
		// dataloader.Loader[string, *model.User]構造体型をセットするために、
		// dataloader.NewBatchedLoader関数を呼び出す
		UserLoader: dataloader.NewBatchedLoader(userBatcher.BatchGetUsers),
	}
}

type userBatcher struct {
	Srv services.Services
}

func (u *userBatcher) BatchGetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// dataloader.Keysの型を[]stringに変換
	IDs := make([]string, len(keys))
	for idx, key := range keys {
		IDs[idx] = key.String()
	}

	// 引数と戻り値のスライスlenは等しくする
	results := make([]*dataloader.Result, len(IDs))
	for i := range results {
		results[i] = &dataloader.Result{
			Error: errors.New("not found"),
		}
	}

	// 検索条件であるIDが、引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexs := make(map[string]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	// サービス層のメソッドを使い、指定されたIDを持つユーザーを全て取得する
	// (ListUsersByIDメソッド内では、IN句を用いたselect文が実行されている)
	users, err := u.Srv.ListUsersByID(ctx, IDs)

	// 取得結果を、戻り値resultの中の適切な場所に格納する
	for _, user := range users {
		var rsl *dataloader.Result
		if err != nil {
			rsl = &dataloader.Result{
				Error: err,
			}
		} else {
			rsl = &dataloader.Result{
				Data: user,
			}
		}
		results[indexs[user.ID]] = rsl
	}
	return results
}
