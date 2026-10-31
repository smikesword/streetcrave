package mock

import (
        context "context"
        reflect "reflect"
        model "smikesword/streetcravela.com/business/pkg/model"

        gomock "go.uber.org/mock/gomock"
)

// MockbusinessRepository is a mock of businessRepository interface.
type MockbusinessRepository struct {
        ctrl     *gomock.Controller
        recorder *MockbusinessRepositoryMockRecorder
        isgomock struct{}
}

// MockbusinessRepositoryMockRecorder is the mock recorder for MockbusinessRepository.
type MockbusinessRepositoryMockRecorder struct {
        mock *MockbusinessRepository
}

// NewMockbusinessRepository creates a new mock instance.
func NewMockbusinessRepository(ctrl *gomock.Controller) *MockbusinessRepository {
        mock := &MockbusinessRepository{ctrl: ctrl}
        mock.recorder = &MockbusinessRepositoryMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockbusinessRepository) EXPECT() *MockbusinessRepositoryMockRecorder {
        return m.recorder
}

// CreateBusiness mocks base method.
func (m *MockbusinessRepository) CreateBusiness(ctx context.Context, business model.Business) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateBusiness", ctx, business)
        ret0, _ := ret[0].(error)
        return ret0
}

// CreateBusiness indicates an expected call of CreateBusiness.
func (mr *MockbusinessRepositoryMockRecorder) CreateBusiness(ctx, business any) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBusiness", reflect.TypeOf((*MockbusinessRepository)(nil).CreateBusiness), ctx, business)
}

// GetAll mocks base method.
func (m *MockbusinessRepository) GetAll(ctx context.Context) ([]model.Business, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAll", ctx)
        ret0, _ := ret[0].([]model.Business)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockbusinessRepositoryMockRecorder) GetAll(ctx any) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockbusinessRepository)(nil).GetAll), ctx)
}

// GetById mocks base method.
func (m *MockbusinessRepository) GetById(ctx context.Context, businessId string) (*model.Business, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetById", ctx, businessId)
        ret0, _ := ret[0].(*model.Business)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockbusinessRepositoryMockRecorder) GetById(ctx, businessId any) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockbusinessRepository)(nil).GetById), ctx, businessId)
}