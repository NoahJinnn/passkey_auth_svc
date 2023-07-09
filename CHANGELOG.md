# Changelog

## [1.0.3](https://github.com/HelloHQ/hqservice/compare/v1.0.2...v1.0.3) (2023-07-09)


### Bug Fixes

* **github:** Fix directory of docker context ([9745860](https://github.com/HelloHQ/hqservice/commit/9745860265f1312e3bf8e243c5cdd03ca2875636))

## [1.0.2](https://github.com/HelloHQ/hqservice/compare/v1.0.1...v1.0.2) (2023-07-09)


### Bug Fixes

* **github:** Decapitalize repo name ([f989058](https://github.com/HelloHQ/hqservice/commit/f989058b44e11640d441031f6b14042f29fced62))

## 1.0.0 (2023-07-09)


### Features

* **app/asset:** Add asset service ([3a82626](https://github.com/HelloHQ/hqservice/commit/3a826263e47a4e38e4efc835876849fe04f724cc))
* **asset/handlers:** Add Asset CRUD handler API ([25b425c](https://github.com/HelloHQ/hqservice/commit/25b425cf2fea382b720263719d01fbb99101c601))
* **asset/handlers:** Add routes for assete CRUD ([3f3faaf](https://github.com/HelloHQ/hqservice/commit/3f3faafaf5a4c084a714f8c16a2eca4366fb861a))
* **asset:** Remove get all assets api ([4372a04](https://github.com/HelloHQ/hqservice/commit/4372a04613f783a07a4c2bb2099824f74fd67fdb))
* **auth/config_Test:** Add doppler client and update expected result in config test ([e9c5a54](https://github.com/HelloHQ/hqservice/commit/e9c5a546028d4324fb5b76ebe6995be6b1e4694e))
* **dal/asset:** Add asset repository, update Asset model, fix test mock ([8eb7efb](https://github.com/HelloHQ/hqservice/commit/8eb7efb857dbc6bd72c60f7e8180ff7cebfd1526))
* **dal:** Implement eager create repo instance for services's dal ([1c32de8](https://github.com/HelloHQ/hqservice/commit/1c32de876e7b0daf274caa079e41ddf4caa68ab7))
* **db:** Add sqlite ent client ([8a98583](https://github.com/HelloHQ/hqservice/commit/8a98583f8efe83cfc985fe2fee0d797402dd6d3e))
* **email.go:** Fix create email, get primary email apis. Run set primary email test successful ([13d968b](https://github.com/HelloHQ/hqservice/commit/13d968beecd46d5529edf00e744dac1fe8eeae18))
* **ent,asset_test:** Update asset schema, setup table test for create asset, update asset routes ([67f849b](https://github.com/HelloHQ/hqservice/commit/67f849bfa89cf20f05b21678ff8e12c5c6d879a4))
* **ent/schema:** Add sql schema endpoints to generate diff for migration ([08770f2](https://github.com/HelloHQ/hqservice/commit/08770f268f060d9ed8b766cf5f3d98b29d11698f))
* **ent:** Add Asset model ([26125cf](https://github.com/HelloHQ/hqservice/commit/26125cfe6c6ff1893f07a291a7097f0e16f85a9c))
* **ent:** Add finverse session model ([75e56d7](https://github.com/HelloHQ/hqservice/commit/75e56d72e83f9807fed328c1ebe3f7850ef3d128))
* **ent:** Update provider schema ([4238ce0](https://github.com/HelloHQ/hqservice/commit/4238ce099127e5d5b761592082fccd572a1e6b01))
* **ent:** Use int32 for numeric types ([fa3b053](https://github.com/HelloHQ/hqservice/commit/fa3b05343d531c213c39a4e90c7e800b3ce48c4c))
* **finverse/auth:** Save session token into postgresql ([e698f58](https://github.com/HelloHQ/hqservice/commit/e698f58edd30353cb1207cd892321ab07b001692))
* **fv/data:** Add api get all accounts, log marshall error ([5c59c77](https://github.com/HelloHQ/hqservice/commit/5c59c773ec70041bcb646a0ce91f07262fe64f02))
* **fv/data:** Complete get all transactions and balance history by account id api ([7897ba2](https://github.com/HelloHQ/hqservice/commit/7897ba290d438acb694187fe2ee963aa9cfeb4ac))
* **fv/handlers:** Add get all institutions API ([fe56b53](https://github.com/HelloHQ/hqservice/commit/fe56b53e59afbc22acbeec7ba758add6f696c7b1))
* **internal/db/sqlite:** Separate ent generation for sqlite, create ent_gen script ([92c8f3f](https://github.com/HelloHQ/hqservice/commit/92c8f3f18eee6a60bbfa6173dcd90cde27ff11b5))
* **internal/session:** Add audience and issuer for JWT session manager ([58bdf82](https://github.com/HelloHQ/hqservice/commit/58bdf829b82f502e28fb744ed7866d46f3e51ced))
* **networth/app:** Add finverse data service for data api ([00fed2e](https://github.com/HelloHQ/hqservice/commit/00fed2eb239f6c9ffcc3bea22c0900ce98bf7343))
* **networth/app:** Init provider service ([8d31f67](https://github.com/HelloHQ/hqservice/commit/8d31f67023a634c1615539cc382b671e0988ccdb))
* **networth/dal:** Add finverse session repo (ongoing refactor dal interface) ([072376c](https://github.com/HelloHQ/hqservice/commit/072376ce5af960c2f204c6a15ee76c2a0b37d601))
* **provider:** Add get connections per user ([dbe7aea](https://github.com/HelloHQ/hqservice/commit/dbe7aea3c1ae25a50ea013e6b85c95010c39f591))
* **provider:** Add save connection ([5076cdb](https://github.com/HelloHQ/hqservice/commit/5076cdbc4952d94112d1ac89978788fc4aa2dcc9))
* **provider:** Create map contains sqlite connection per user ([5b39edd](https://github.com/HelloHQ/hqservice/commit/5b39edd0303d1237fff849dd4a0bde21c0caab7c))
* **router:** Add session middleware for fv apis ([1885601](https://github.com/HelloHQ/hqservice/commit/1885601462861f8b64717b9b6d303de2ca5fb6c2))
* **scripts/sql:** Create taskfile script to run atlas migration generation ([7a80ae9](https://github.com/HelloHQ/hqservice/commit/7a80ae9c397e70cd6a91dc8c8a36919bd3fa57ce))
* **scripts:** Add migrate script using atlas ([a588022](https://github.com/HelloHQ/hqservice/commit/a5880223821626f823a54c53a01cf34286b76104))
* **scripts:** Create apply migration scripts ([ee3f6f3](https://github.com/HelloHQ/hqservice/commit/ee3f6f3cb7082a1f07191b927bf7afe7f2bf4aba))
* **scripts:** Create migration scripts ([74b61cd](https://github.com/HelloHQ/hqservice/commit/74b61cdb8334c695e3b34ee73d47ef46204266ff))
* **sqlite/ent:** Create Institution, Connection, Account, Assets table ([c1d0f7e](https://github.com/HelloHQ/hqservice/commit/c1d0f7ef7405a0f21461c2feb36ebeaaf4f1c78b))
