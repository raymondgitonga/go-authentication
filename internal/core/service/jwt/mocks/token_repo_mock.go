// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package jwt_mocks

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/raymondgitonga/go-authentication/internal/core/service/jwt"
	"sync"
)

// Ensure, that TokenRepositoryMock does implement TokenRepository.
// If this is not the case, regenerate this file with moq.
var _ jwt.TokenRepository = &TokenRepositoryMock{}

// TokenRepositoryMock is a mock implementation of TokenRepository.
//
//	func TestSomethingThatUsesTokenRepository(t *testing.T) {
//
//		// make and configure a mocked TokenRepository
//		mockedTokenRepository := &TokenRepositoryMock{
//			GetLatestTokenFunc: func(ctx context.Context) (*redis.Z, error) {
//				panic("mock out the GetLatestToken method")
//			},
//			GetTokenFunc: func(ctx context.Context, tokenID int64) (*redis.Z, error) {
//				panic("mock out the GetToken method")
//			},
//		}
//
//		// use mockedTokenRepository in code that requires TokenRepository
//		// and then make assertions.
//
//	}
type TokenRepositoryMock struct {
	// GetLatestTokenFunc mocks the GetLatestToken method.
	GetLatestTokenFunc func(ctx context.Context) (*redis.Z, error)

	// GetTokenFunc mocks the GetToken method.
	GetTokenFunc func(ctx context.Context, tokenID int64) (*redis.Z, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetLatestToken holds details about calls to the GetLatestToken method.
		GetLatestToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetToken holds details about calls to the GetToken method.
		GetToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// TokenID is the tokenID argument value.
			TokenID int64
		}
	}
	lockGetLatestToken sync.RWMutex
	lockGetToken       sync.RWMutex
}

// GetLatestToken calls GetLatestTokenFunc.
func (mock *TokenRepositoryMock) GetLatestToken(ctx context.Context) (*redis.Z, error) {
	if mock.GetLatestTokenFunc == nil {
		panic("TokenRepositoryMock.GetLatestTokenFunc: method is nil but TokenRepository.GetLatestToken was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetLatestToken.Lock()
	mock.calls.GetLatestToken = append(mock.calls.GetLatestToken, callInfo)
	mock.lockGetLatestToken.Unlock()
	return mock.GetLatestTokenFunc(ctx)
}

// GetLatestTokenCalls gets all the calls that were made to GetLatestToken.
// Check the length with:
//
//	len(mockedTokenRepository.GetLatestTokenCalls())
func (mock *TokenRepositoryMock) GetLatestTokenCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetLatestToken.RLock()
	calls = mock.calls.GetLatestToken
	mock.lockGetLatestToken.RUnlock()
	return calls
}

// GetToken calls GetTokenFunc.
func (mock *TokenRepositoryMock) GetToken(ctx context.Context, tokenID int64) (*redis.Z, error) {
	if mock.GetTokenFunc == nil {
		panic("TokenRepositoryMock.GetTokenFunc: method is nil but TokenRepository.GetToken was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		TokenID int64
	}{
		Ctx:     ctx,
		TokenID: tokenID,
	}
	mock.lockGetToken.Lock()
	mock.calls.GetToken = append(mock.calls.GetToken, callInfo)
	mock.lockGetToken.Unlock()
	return mock.GetTokenFunc(ctx, tokenID)
}

// GetTokenCalls gets all the calls that were made to GetToken.
// Check the length with:
//
//	len(mockedTokenRepository.GetTokenCalls())
func (mock *TokenRepositoryMock) GetTokenCalls() []struct {
	Ctx     context.Context
	TokenID int64
} {
	var calls []struct {
		Ctx     context.Context
		TokenID int64
	}
	mock.lockGetToken.RLock()
	calls = mock.calls.GetToken
	mock.lockGetToken.RUnlock()
	return calls
}
