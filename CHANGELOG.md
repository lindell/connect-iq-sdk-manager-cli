# Changelog

## [0.4.0](https://github.com/lindell/connect-iq-sdk-manager-cli/compare/v0.3.1...v0.4.0) (2023-08-16)


### Features

* added "sdk list" command ([dacfc1a](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/dacfc1a398d9c32bbfe1f40c4930d710dc686a94))
* added configurable logging ([c7171b9](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/c7171b9c975079f49958aa9ae0bc4cdc8168aca5))
* added list device command and moved it and device download under ([0d1d42a](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/0d1d42a9cf252312735991ad777b603d5d7b02ee))
* added sdk current-path command ([3bac92b](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/3bac92bab66be6ccc6127aadd3028634291498d1))
* added sdk download command ([d79508b](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/d79508b057b2013b6644f83acc49ec7172c545e0))
* added sdk set command ([41a737d](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/41a737db0510dd69961e37a167115914ce83bf0b))
* handle token refresh ([#16](https://github.com/lindell/connect-iq-sdk-manager-cli/issues/16)) ([781b2f7](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/781b2f7634c81f45711517fc47a8069fe145f91b))
* replaced asking for username and password with real oauth flow ([#18](https://github.com/lindell/connect-iq-sdk-manager-cli/issues/18)) ([5108950](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/51089500b035eff44848c729e47823a84c27f0c3))


### Bug Fixes

* don't log downloaded file contents ([2e64ff6](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/2e64ff6a5092a2da5ecaf01f39fd2d76661416e1))
* improved error message when login fails ([fc55b2f](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/fc55b2f57e76b1fb2a6494aff5066ecae61cf54b)), closes [#2](https://github.com/lindell/connect-iq-sdk-manager-cli/issues/2)

## [0.3.1](https://github.com/lindell/connect-iq-sdk-manager-cli/compare/v0.3.0...v0.3.1) (2023-08-10)


### Bug Fixes

* changed binary name ([9200a0a](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/9200a0a53126906c210a8e3fc1f5d439a84f419a))
* fixed install script project names ([d20b38d](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/d20b38d62a7a210b5c2c26e8ddcf46f7235b6729))

## [0.3.0](https://github.com/lindell/connect-iq-sdk-manager-cli/compare/v0.2.1...v0.3.0) (2023-08-09)


### Features

* renamed to connect-iq-sdk-manager(-cli) ([0c05f57](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/0c05f574c0cb445be191ec4306f71ef845c292e7))


### Bug Fixes

* remove duplicate devices before further processing ([c98632f](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/c98632f8d8f393d9a66fbfabc8b0fb8f618a9f1f))

## [0.2.1](https://github.com/lindell/connect-iq-sdk-manager-cli/compare/v0.2.0...v0.2.1) (2023-08-03)


### Bug Fixes

* fixed permission of garmin folder ([e57d221](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/e57d221cd3afd3f177e7b22690a4a36f4557b88a))

## [0.2.0](https://github.com/lindell/connect-iq-sdk-manager-cli/compare/v0.1.0...v0.2.0) (2023-08-03)


### Features

* added option to set login credentials via env vars ([ba8cddb](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/ba8cddba3c0c6105362fca21fed3a73cb0143a62))


### Bug Fixes

* fixed old cli description ([fd9b2b1](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/fd9b2b1f1cf5891bca6d34f290c7f2271678e1cf))

## 0.1.0 (2023-08-03)


### Features

* added basic download command ([#5](https://github.com/lindell/connect-iq-sdk-manager-cli/issues/5)) ([1a5ed1e](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/1a5ed1e7b975706c2719ca5666a9df28dda552d8))
* added login command ([516b5d9](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/516b5d9a30b9f680d82a3c33072cad6253ba0fae))


### Bug Fixes

* fixed windows password reading ([e247f44](https://github.com/lindell/connect-iq-sdk-manager-cli/commit/e247f44807a892aedc0b40f95ac3a21ba6d42b64))
