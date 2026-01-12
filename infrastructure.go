// Copyright 2025 zampo.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// @contact  zampo3380@gmail.com

package bootstrap

import (
	"github.com/go-anyway/framework-core"
	"github.com/go-anyway/framework-log"

	"go.uber.org/zap"
)

// InitInfrastructureForServer 初始化 RPC 服务的基础设施（仅日志）
// 注意：trace 的初始化由 AppBuilder.Build() 统一管理（包含初始化和生命周期关闭）
// 这样设计是为了：
// 1. 避免重复初始化 trace
// 2. 确保 trace 的关闭由 App 的生命周期管理，保证优雅关闭
// registry: 配置注册表（从配置中获取日志配置进行初始化）
func InitInfrastructureForServer(registry *app.DefaultConfigRegistry) {
	// 获取日志配置
	logCfg, err := registry.GetLog()
	if err != nil {
		log.Fatal("failed to get log config", zap.Error(err))
	}
	// 初始化日志
	initLogger(logCfg)
	// trace 初始化由 AppBuilder.Build() 统一处理，这里不处理
}

// initLogger 初始化日志（用于 RPC 服务）
func initLogger(cfg *log.Config) {
	logOpts := cfg.ToOptions()

	// 初始化日志
	log.Init(
		log.WithLevel(logOpts.Level),
		log.WithFormat(logOpts.Format),
		log.WithOutputPaths(logOpts.OutputPaths),
		log.WithErrorOutputPaths(logOpts.ErrorOutputPaths),
		log.WithDisableCaller(logOpts.DisableCaller),
		log.WithDisableStacktrace(logOpts.DisableStacktrace),
		log.WithFilename(logOpts.Filename),
		log.WithMaxSize(logOpts.MaxSize),
		log.WithMaxAge(logOpts.MaxAge),
		log.WithMaxBackups(logOpts.MaxBackups),
		log.WithCompress(logOpts.Compress),
		log.WithDevelopment(logOpts.Development),
	)
}
