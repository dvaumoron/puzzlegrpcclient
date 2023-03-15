/*
 *
 * Copyright 2023 puzzlegrpcclient authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package puzzlegrpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type Client struct {
	serviceAddr string
	dialOptions grpc.DialOption
	timeOut     time.Duration
}

func Make(serviceAddr string, dialOptions grpc.DialOption, timeOut time.Duration) Client {
	return Client{serviceAddr: serviceAddr, dialOptions: dialOptions, timeOut: timeOut}
}

func (client Client) Dial() (*grpc.ClientConn, error) {
	return grpc.Dial(client.serviceAddr, client.dialOptions)
}

func (client Client) InitContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), client.timeOut)
}
