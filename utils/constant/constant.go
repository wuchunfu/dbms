/*
Copyright © 2020 Marvin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package constant

// Database Type
const (
	DatabaseTypeOracle = "ORACLE"
	DatabaseTypeMySQL  = "MYSQL"
	DatabaseTypeTiDB   = "TiDB"
)

// DefaultRecordCreateBatchSize Model Create Record Default Batch Size
const DefaultRecordCreateBatchSize = 50

// StringSeparatorComma used for build-in string data separator
const StringSeparatorComma = ","
const StringSeparatorSlash = "/"

const (
	MigrateTaskTableIsExclude    = "YES"
	MigrateTaskTableIsNotExclude = "NO"
)

// DefaultTaskQueueChannelSize used for task queue channel size
const DefaultTaskQueueChannelSize = 1024

// DefaultDataEncryptDecryptKey used for data encrypt and decrypt key
const DefaultDataEncryptDecryptKey = "marvin@jwt!#$&123"
