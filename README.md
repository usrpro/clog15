[![Build Status](https://travis-ci.org/usrpro/clog15.svg?branch=master)](https://travis-ci.org/usrpro/clog15)
[![codecov](https://codecov.io/gh/usrpro/clog15/branch/master/graph/badge.svg)](https://codecov.io/gh/usrpro/clog15)

# Clog15

Package clog15 provides utilities to embed and extract a log15.Logger to a context.
This might be helpfull to preserve logger context while being restricted by funtction signatures.
For instance in http.HanderFunc, middleware or gRPC interceptors.
It allows you to define logging context and attach the configured logger to a context
passed down the executions chain.

## License
Copyright (c) 2020, Mohlmann Solutions SRL. All rights reserved.
Use of this source code is governed by a BSD 3 Clause License that can be found in the [LICENSE](LICENSE) file.
