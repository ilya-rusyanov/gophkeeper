package store

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/store/mock"
)

func TestStoreBin(t *testing.T) {
	ctx := context.Background()

	t.Run("successfull store", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		authRepo := mock.NewMockAuthStorager(ctrl)
		service := mock.NewMockServicer(ctrl)
		fileRead := mock.NewMockFileReader(ctrl)

		authRepo.EXPECT().
			Load().
			Return(entity.NewMyAuthentication("auth"), nil)

		bin, err := hex.DecodeString("ffd8ffe000104a46")
		require.NoError(t, err)

		fileRead.EXPECT().
			ReadFile(gomock.Any(), "/tmp/view.jpeg").
			Return(bin, nil)

		service.EXPECT().
			Store(gomock.Any(),
				*entity.NewServiceStoreRequest(
					entity.NewMyAuthentication(
						"auth",
					),
					*entity.NewBinRecord(
						"img",
						entity.Meta{"theme:sea"},
						entity.BinPayload(bin),
					),
				),
			)

		uc := NewBin(authRepo, fileRead, service)

		err = uc.StoreBin(ctx,
			*entity.NewBinRecord(
				"img",
				entity.Meta{"theme:sea"},
				[]byte{},
			),
			"/tmp/view.jpeg",
		)

		assert.NoError(t, err)
	})
}
