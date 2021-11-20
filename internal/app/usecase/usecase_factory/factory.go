package usecase_factory

import (
	usecase_csrf "patreon/internal/app/csrf/usecase"
	useAccess "patreon/internal/app/usecase/access"
	useAwards "patreon/internal/app/usecase/awards"
	useCreator "patreon/internal/app/usecase/creator"
	useInfo "patreon/internal/app/usecase/info"
	useLikes "patreon/internal/app/usecase/likes"
	usePayments "patreon/internal/app/usecase/payments"
	usePosts "patreon/internal/app/usecase/posts"
	useAttaches "patreon/internal/app/usecase/attaches"
	useSubscr "patreon/internal/app/usecase/subscribers"
	useUser "patreon/internal/app/usecase/user"
	"patreon/internal/microservices/auth/sessions"
)

type UsecaseFactory struct {
	repositoryFactory  RepositoryFactory
	userUsecase        useUser.Usecase
	creatorUsecase     useCreator.Usecase
	csrfUsecase        usecase_csrf.Usecase
	accessUsecase      useAccess.Usecase
	subscribersUsecase useSubscr.Usecase
	awardsUsercase     useAwards.Usecase
	awardsUsecase      useAwards.Usecase
	sessinManager      sessions.SessionsManager
	postsUsecase       usePosts.Usecase
	attachesUsecase   useAttaches.Usecase
	infoUsecase        useInfo.Usecase
	likesUsecase       useLikes.Usecase
	paymentsUsecase    usePayments.Usecase
}

func NewUsecaseFactory(repositoryFactory RepositoryFactory) *UsecaseFactory {
	return &UsecaseFactory{
		repositoryFactory: repositoryFactory,
	}
}

func (f *UsecaseFactory) GetUserUsecase() useUser.Usecase {
	if f.userUsecase == nil {
		f.userUsecase = useUser.NewUserUsecase(f.repositoryFactory.GetUserRepository(), f.repositoryFactory.GetFilesRepository())
	}
	return f.userUsecase
}

func (f *UsecaseFactory) GetCreatorUsecase() useCreator.Usecase {
	if f.creatorUsecase == nil {
		f.creatorUsecase = useCreator.NewCreatorUsecase(f.repositoryFactory.GetCreatorRepository(),
			f.repositoryFactory.GetFilesRepository())
	}
	return f.creatorUsecase
}
func (f *UsecaseFactory) GetCsrfUsecase() usecase_csrf.Usecase {
	if f.csrfUsecase == nil {
		f.csrfUsecase = usecase_csrf.NewCsrfUsecase(f.repositoryFactory.GetCsrfRepository())
	}
	return f.csrfUsecase
}

func (f *UsecaseFactory) GetAccessUsecase() useAccess.Usecase {
	if f.accessUsecase == nil {
		f.accessUsecase = useAccess.NewAccessUsecase(f.repositoryFactory.GetAccessRepository())
	}
	return f.accessUsecase
}
func (f *UsecaseFactory) GetSubscribersUsecase() useSubscr.Usecase {
	if f.subscribersUsecase == nil {
		f.subscribersUsecase = useSubscr.NewSubscribersUsecase(f.repositoryFactory.GetSubscribersRepository(),
			f.repositoryFactory.GetAwardsRepository())
	}
	return f.subscribersUsecase
}

func (f *UsecaseFactory) GetAwardsUsecase() useAwards.Usecase {
	if f.awardsUsecase == nil {
		f.awardsUsecase = useAwards.NewAwardsUsecase(f.repositoryFactory.GetAwardsRepository(),
			f.repositoryFactory.GetFilesRepository())
	}
	return f.awardsUsecase
}

func (f *UsecaseFactory) GetPostsUsecase() usePosts.Usecase {
	if f.postsUsecase == nil {
		f.postsUsecase = usePosts.NewPostsUsecase(f.repositoryFactory.GetPostsRepository(),
			f.repositoryFactory.GetAttachesRepository(), f.repositoryFactory.GetFilesRepository())
	}
	return f.postsUsecase
}

func (f *UsecaseFactory) GetLikesUsecase() useLikes.Usecase {
	if f.likesUsecase == nil {
		f.likesUsecase = useLikes.NewLikesUsecase(f.repositoryFactory.GetLikesRepository())
	}
	return f.likesUsecase
}

func (f *UsecaseFactory) GetAttachesUsecase() useAttaches.Usecase {
	if f.attachesUsecase == nil {
		f.attachesUsecase = useAttaches.NewAttachesUsecase(f.repositoryFactory.GetAttachesRepository(),
			f.repositoryFactory.GetFilesRepository())
	}
	return f.attachesUsecase
}

func (f *UsecaseFactory) GetPaymentsUsecase() usePayments.Usecase {
	if f.paymentsUsecase == nil {
		f.paymentsUsecase = usePayments.NewPaymentsUsecase(f.repositoryFactory.GetPaymentsRepository())
	}
	return f.paymentsUsecase
}

func (f *UsecaseFactory) GetInfoUsecase() useInfo.Usecase {
	if f.infoUsecase == nil {
		f.infoUsecase = useInfo.NewInfoUsecase(f.repositoryFactory.GetInfoRepository())
	}
	return f.infoUsecase
}
