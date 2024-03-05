package show

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	"github.com/ilya-rusyanov/gophkeeper/internal/client/usecase/show/mock"
)

//go:generate mockgen -destination ./mock/mocks.go -package mock . Servicer,Storager,FileSaver

func TestShow(t *testing.T) {
	ctx := context.Background()

	t.Run("ordinary show", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		storager := mock.NewMockStorager(ctrl)
		servicer := mock.NewMockServicer(ctrl)

		storager.EXPECT().
			Load().
			Return(entity.NewMyAuthentication("my auth"), nil)

		servicer.EXPECT().
			Show(gomock.Any(),
				entity.ServiceShowRequest{
					AuthData: entity.NewMyAuthentication("my auth"),
					Type:     entity.RecordTypeAuth,
					Name:     "yandex",
				}).Return(
			*entity.NewAuthRecord(
				"yandex",
				entity.Meta{"expires:july"},
				entity.NewAuthPayload("elon", "twitterx"),
			), nil)

		uc := New(storager, servicer, nil)

		s, err := uc.Show(ctx, entity.ShowIn{
			Type: "auth",
			Name: "yandex",
		})

		require.NoError(t, err)
		assert.Equal(t, *entity.NewAuthRecord(
			"yandex",
			entity.Meta{"expires:july"},
			entity.NewAuthPayload("elon", "twitterx"),
		), s)
	})

	t.Run("binary show", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		storager := mock.NewMockStorager(ctrl)
		servicer := mock.NewMockServicer(ctrl)
		fileSaver := mock.NewMockFileSaver(ctrl)

		storager.EXPECT().
			Load().
			Return(entity.NewMyAuthentication("my auth"), nil)

		bin, err := hex.DecodeString("ffd8ffe000104a46")
		require.NoError(t, err)

		binPayload := entity.BinPayload(bin)

		servicer.EXPECT().
			Show(gomock.Any(),
				entity.ServiceShowRequest{
					AuthData: entity.NewMyAuthentication("my auth"),
					Type:     entity.RecordTypeBin,
					Name:     "img",
				}).Return(
			*entity.NewBinRecord(
				"img",
				entity.Meta{"theme:sea"},
				binPayload,
			), nil)

		fileSaver.EXPECT().
			SaveFile(gomock.Any(), entity.FileSaveIn{
				Data:     &binPayload,
				FilePath: "/tmp/image.jpeg",
			})

		uc := New(storager, servicer, fileSaver)

		err = uc.ShowBin(ctx, entity.ShowBinIn{
			Name:   "img",
			SaveTo: "/tmp/image.jpeg",
		})
		require.NoError(t, err)
	})
}
