package scanner

import (
	"context"

	"github.com/guneyin/gobist"
	"github.com/guneyin/kapscan/entity"
	"github.com/guneyin/kapscan/util"
)

type Repo struct {
	bist *gobist.Bist
}

func NewRepo() *Repo {
	return &Repo{bist: gobist.New()}
}

func (r *Repo) GetSymbolList(ctx context.Context) (entity.CompanyList, error) {
	symbolList, err := r.bist.GetSymbolList(ctx)
	if err != nil {
		return nil, err
	}

	cl := make(entity.CompanyList, symbolList.Count)
	for i, symbol := range symbolList.Items {
		cl[i] = entity.Company{
			Code: symbol.Code,
			Name: symbol.Name,
			Icon: symbol.Icon,
		}
	}

	return cl, nil
}

func (r *Repo) SyncCompanyWithShares(ctx context.Context, cmp *entity.Company) error {
	fetched, err := r.bist.GetCompanyWithShares(ctx, cmp.Code)
	if err != nil {
		return err
	}

	_, err = util.Convert(fetched, cmp)
	if err != nil {
		return err
	}

	return nil
}
