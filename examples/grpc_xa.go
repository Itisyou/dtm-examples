/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/dtm/dtmcli/logger"
	"github.com/dtm-labs/dtm/dtmgrpc"
	"github.com/dtm-labs/dtm/dtmsvr"
	"github.com/dtm-labs/dtm/test/busi"
	"google.golang.org/protobuf/types/known/emptypb"
)

func init() {
	addSample("grpc_xa", func() string {
		gid := dtmgrpc.MustGenGid(dtmsvr.DefaultGrpcServer)
		req := &busi.BusiReq{Amount: 30}
		err := busi.XaGrpcClient.XaGlobalTransaction(gid, func(xa *dtmgrpc.XaGrpc) error {
			r := &emptypb.Empty{}
			err := xa.CallBranch(req, busi.BusiGrpc+"/examples.Busi/TransOutXa", r)
			if err != nil {
				return err
			}
			err = xa.CallBranch(req, busi.BusiGrpc+"/examples.Busi/TransInXa", r)
			return err
		})
		logger.FatalIfError(err)
		return gid
	})
}
