package userservise

import (
	"Q/A-GameApp/dto"
	"Q/A-GameApp/pkg/richerror"
)

// All Requests inputs for interactor / service should sanitize

func (s Service) Profile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	const op = "userService.profile"
	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return dto.ProfileResponse{}, richerror.New(op).
			WhitWarpError(err).
			WhitMeta(map[string]interface{}{"request": req})
	}
	return dto.ProfileResponse{Name: user.Name}, nil
}
