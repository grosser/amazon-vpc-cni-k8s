// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package retry

type Retriable interface {
	Retry() bool
}

type DefaultRetriable struct {
	retry bool
}

func (dr DefaultRetriable) Retry() bool {
	return dr.retry
}

func NewRetriable(retry bool) Retriable {
	return DefaultRetriable{
		retry: retry,
	}
}

type RetriableError interface {
	Retriable
	error
}

type DefaultRetriableError struct {
	Retriable
	error
}

func NewRetriableError(retriable Retriable, err error) RetriableError {
	return &DefaultRetriableError{
		retriable,
		err,
	}
}
