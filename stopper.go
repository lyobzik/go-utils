//Copyright 2016 lyobzik
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package utils

import (
	"sync"
)

type Stopper struct {
	waitDone sync.WaitGroup
	Stopping chan struct{}
}

func NewStopper() *Stopper {
	return &Stopper{
		waitDone: sync.WaitGroup{},
		Stopping: make(chan struct{}, 1),
	}
}

func (s *Stopper) Stop() {
	close(s.Stopping)
}

func (s *Stopper) WaitDone() {
	s.waitDone.Wait()
}

func (s *Stopper) Add() {
	s.waitDone.Add(1)
}

func (s *Stopper) Done() {
	s.waitDone.Done()
}

func (s *Stopper) Join() {
	s.Stop()
	s.WaitDone()
}
