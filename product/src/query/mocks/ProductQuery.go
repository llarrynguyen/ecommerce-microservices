// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import query "github.com/wuriyanto48/ecommerce-grpc-microservice/product/src/query"

// ProductQuery is an autogenerated mock type for the ProductQuery type
type ProductQuery struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *ProductQuery) FindAll() <-chan query.QueryResult {
	ret := _m.Called()

	var r0 <-chan query.QueryResult
	if rf, ok := ret.Get(0).(func() <-chan query.QueryResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan query.QueryResult)
		}
	}

	return r0
}

// FindByCategory provides a mock function with given fields: categoryID
func (_m *ProductQuery) FindByCategory(categoryID int) <-chan query.QueryResult {
	ret := _m.Called(categoryID)

	var r0 <-chan query.QueryResult
	if rf, ok := ret.Get(0).(func(int) <-chan query.QueryResult); ok {
		r0 = rf(categoryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan query.QueryResult)
		}
	}

	return r0
}

// FindByID provides a mock function with given fields: id
func (_m *ProductQuery) FindByID(id int) <-chan query.QueryResult {
	ret := _m.Called(id)

	var r0 <-chan query.QueryResult
	if rf, ok := ret.Get(0).(func(int) <-chan query.QueryResult); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan query.QueryResult)
		}
	}

	return r0
}