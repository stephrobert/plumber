## 1.0.0 (2026-03-05)


### ✨ Features

* **analysis:** Add --mr-comment to create mr comments during analysis. Add --badge to create/update Plumber compliance badge when running on default remote branch ([4cba483](https://github.com/stephrobert/plumber/commit/4cba4839ce3512ef2da7ff5b9512dca701d2a7f7))
* **analysis:** Allow auto-detection for gitlab url and project during analysis + update banner ([e7a20e6](https://github.com/stephrobert/plumber/commit/e7a20e6e2b49bdf7cc7805effebc765e73930b2d))
* **analysis:** Revert CI_JOB_TOKEN ([6c12fb5](https://github.com/stephrobert/plumber/commit/6c12fb59ad73147ca96c3e58ce570eab751706eb))
* **analyze:** add --controls and --skip-controls control filtering ([9a9aca0](https://github.com/stephrobert/plumber/commit/9a9aca0a50e63eb27a7f7c8f5470ae6922ff425c))
* **analyze:** Make conf and threshold optional ([bf6a4df](https://github.com/stephrobert/plumber/commit/bf6a4dfc3c306d0bb9bd3cd8e7e6a20bcaf9522a))
* **artifact:** Add new concept: Pipeline Bill Of Materials (PBOM) and add cyclonedx output format support ([7097605](https://github.com/stephrobert/plumber/commit/7097605f85123ea7599e67d7b22e34bfa13e726b))
* **brew:** Enable automatic updating of brew tap formula repo upon new release ([ead9860](https://github.com/stephrobert/plumber/commit/ead98601a2e355403975b087a5f0da1b3653fb76))
* **brew:** Test release 0.1.26 ([1848d52](https://github.com/stephrobert/plumber/commit/1848d5248fd9c38cc54d8da11b24ef9694afb48d))
* **build:** Add platforms binary releases ([01d9bfa](https://github.com/stephrobert/plumber/commit/01d9bfa79b80cf9d2c7ea6d1984e2acc7a4db9bb))
* **build:** Move to alpine to make command customizable in CI ([763bcf3](https://github.com/stephrobert/plumber/commit/763bcf3eadd21fdf53503033b00ced44b1a6b862))
* **ci:** Add dependabot ([3480a77](https://github.com/stephrobert/plumber/commit/3480a777b3ef3bcf1dd72269da67aad28c3fcf0b))
* **ci:** Replace trivvy with grype ([984286f](https://github.com/stephrobert/plumber/commit/984286f2ca871792b8a12e34bea7a647cb72450f))
* **ci:** Run on ubuntu 24.04 instead of latest ([b5473d2](https://github.com/stephrobert/plumber/commit/b5473d2d527dc6672f332670d729f2eb4701944f))
* **ci:** Start adding test, lint, scan and pin versions by digest ([ec77e70](https://github.com/stephrobert/plumber/commit/ec77e709864ddf621ca5337de3345c5c18627466))
* **cmd:** Add --fail-warnings on the analyze and config validate commands ([fbc5839](https://github.com/stephrobert/plumber/commit/fbc58391ebc1e51d412588ed5748587a6e125227))
* **cmd:** Add progress spinner during analysis ([80b1bad](https://github.com/stephrobert/plumber/commit/80b1bad2a9b556d0e8414e2c90b4199cda5b26c4)), closes [#64](https://github.com/stephrobert/plumber/issues/64)
* **cmd:** notify user when a newer version of plumber is available ([deca33f](https://github.com/stephrobert/plumber/commit/deca33f74d9c2f1c2468d8ce98a5693df1fa9d9c)), closes [#39](https://github.com/stephrobert/plumber/issues/39)
* **cmd:** validating config file before analysis ([a1dc4bd](https://github.com/stephrobert/plumber/commit/a1dc4bd7c14d4fe230f770849c0b12683fba8323))
* **component:** Allow verbosity in component ([b59015a](https://github.com/stephrobert/plumber/commit/b59015a9bd64213ca0b3cf21ea02b16e833eee13))
* **conf:** Add conf view and move generate under conf ([8e549e9](https://github.com/stephrobert/plumber/commit/8e549e97ab462dbf46d7f5d25dea8fd77989c796))
* **conf:** Add reference to examples in test file for required includes ([a8ec829](https://github.com/stephrobert/plumber/commit/a8ec82961ab5152a3b6a1bff938afbc63978b7e6))
* **conf:** Allow conf generation with command ([7390e76](https://github.com/stephrobert/plumber/commit/7390e763055c23a44506ba43db2331fb09a0b0e7))
* **conf:** Correct dockerfile and release file ([34263ec](https://github.com/stephrobert/plumber/commit/34263ec71a90183a65e9b1511e90f34ce6d2e1a2))
* **config:** Add ValidateKnownKeys to warn on unknown config keys ([4c33ca3](https://github.com/stephrobert/plumber/commit/4c33ca3c980cc57419b7b457bfae09c7543ba8d3)), closes [#58](https://github.com/stephrobert/plumber/issues/58) [#58](https://github.com/stephrobert/plumber/issues/58)
* **conf:** Introduce priority and automatic detection of conf files ([91ef31b](https://github.com/stephrobert/plumber/commit/91ef31b3a43b1fa150a8fba1c2c880a2220b80ef))
* **control:** Add 3 new controls ([591a850](https://github.com/stephrobert/plumber/commit/591a8509f47c1cc21eb3bf71ad185163368ba033))
* **control:** Make component collecetion compatible with gitlab built-in components ([532f071](https://github.com/stephrobert/plumber/commit/532f071544c58f8c3af1cbf4771b43a1e296a799))
* **controls:** Add debug trace detection control ([5e08b97](https://github.com/stephrobert/plumber/commit/5e08b974d17ee2468768f6b8a2917b09eddfe0ff)), closes [#86](https://github.com/stephrobert/plumber/issues/86)
* **controls:** add image digest pinning control ([ea538a9](https://github.com/stephrobert/plumber/commit/ea538a954def56d4695b0d766f3bb1aff8ee7bbd))
* **controls:** Add overridden component and templates issue & integration into pbom and cyclonedex ([e970c27](https://github.com/stephrobert/plumber/commit/e970c27697c29873057a5f6b6109715ba47d3896))
* **controls:** Add unsafe variable expansion control for user-filled predefined variables ([0f7ec72](https://github.com/stephrobert/plumber/commit/0f7ec723ff0b03ffb1c345fa393eaa8f9e022dd7))
* **controls:** Rename control outputs and config to make them more human-readable & Start using CI_JOB_TOKEN if in the CI ([6669707](https://github.com/stephrobert/plumber/commit/66697073a13a337230f88a5ba1213def645f474d))
* **control:** Support Natural Language in pipeline inclusion for templates and components ([59c4edd](https://github.com/stephrobert/plumber/commit/59c4edddb2750296d06d282e7c5efd272e3aa81d))
* **init:** Add first version of plumber with 3 controls, gitlab-ci and in-terminal support ([278c2d9](https://github.com/stephrobert/plumber/commit/278c2d9ca38b4ac434ff49587abde53c9eb825f6))
* **license:** Update license in readme to MPL-2.0 ([4cbab86](https://github.com/stephrobert/plumber/commit/4cbab86370f2ee3c5d010d5888e17e1d6fa9d445))
* **local:** Enable lint, validation and analysis of local .gitlab-ci.yml as well as local reoslution of include: local types. ([5a2a3aa](https://github.com/stephrobert/plumber/commit/5a2a3aa3b67c8489289bf98f47f9494f386f6458))
* **log:** Improve logging experience ([426bcf8](https://github.com/stephrobert/plumber/commit/426bcf817bc62c382b5487443832b49372cbe890))
* **naming:** Rename components to plumber, no need for the analyze suffix ([53a0816](https://github.com/stephrobert/plumber/commit/53a08165e7d879099606505c98cce7abc45715eb))
* **output:** Improve readability of printed results ([97d708f](https://github.com/stephrobert/plumber/commit/97d708f75d5b76fba6bb3724543837eb371f5c04))
* **release:** Downgrade feat to patch ([eb30e81](https://github.com/stephrobert/plumber/commit/eb30e8183466068954edbd4e700986e02bfd72af))
* **update:** Empty commit ([b7bd04f](https://github.com/stephrobert/plumber/commit/b7bd04fa8ab8430c15afd625ea27ce3ae2030e8e))
* **UX:** Define default output file, add output json example ([3dbfa1c](https://github.com/stephrobert/plumber/commit/3dbfa1c9172b561e35ce52d834e794bc2fb08091))
* **UX:** If a control is misisng from .plumber.yaml simply skip it instead of returning an error ([3eec388](https://github.com/stephrobert/plumber/commit/3eec388d75d0e28792c580dc95c79f89663af5e3))
* **UX:** Integrate the control pinned by digest inside the immutable one ([87bd450](https://github.com/stephrobert/plumber/commit/87bd45074c7d0710f67dfe23e486421eda1baf39))
* **version:** async update check with opt-out ([a1ef745](https://github.com/stephrobert/plumber/commit/a1ef74553e85c2341e803ea4ec683f6f0c4d6e31))


### 🐛 Bug Fixes

* **analysis:** Fix bug where analyzed branch was being mistaken for branches to protect ([afdd5f8](https://github.com/stephrobert/plumber/commit/afdd5f8424a3a05120e50feb3dd99909bd5dba6a))
* **analysis:** If no controls ran (e.g., data collection failed), compliance is 0% - we can't verify anything ([7ec0e72](https://github.com/stephrobert/plumber/commit/7ec0e72ca459539c4a8e4d4fdbc04faf01ddfae8))
* **branch:** use correct SHA for ciConfig query when --branch is specified ([1729084](https://github.com/stephrobert/plumber/commit/1729084509de141d282e3a3d49c62fe7864e385e))
* **brew:** Typo in release ([6c80575](https://github.com/stephrobert/plumber/commit/6c80575c26f7294e8fa3a1553d10aa3f88864adb))
* **brew:** Typo in release ([4d7b905](https://github.com/stephrobert/plumber/commit/4d7b905d84e47f25cda4c7b770b0f8660f453f7e))
* **bug:** Cleanup some dead code ([fa7e1ae](https://github.com/stephrobert/plumber/commit/fa7e1ae3f445ed2047c5ab4ef0a4a66f0bc8ad93))
* **build:** Move release creation to after asset upload ([bc96e39](https://github.com/stephrobert/plumber/commit/bc96e39f0a4fc56185b62199a1f0da8e6a503367))
* **ci:** Fix CI lint issues ([5288379](https://github.com/stephrobert/plumber/commit/528837904a46b7b6334bbfe5f767d200faebdc81))
* **comment:** Add timeout comment to client ([15df3f0](https://github.com/stephrobert/plumber/commit/15df3f002dea5c9ba5a6a95b9943eafa1dc0230d))
* **component:** Add full docker path to plumber as trusted ([d7732c8](https://github.com/stephrobert/plumber/commit/d7732c8d1fefd7f12730e60e963a703d2a120a64))
* **config:** Fix compilation issues + make validation recursive to test subkeys ([405abe4](https://github.com/stephrobert/plumber/commit/405abe4dc7f78f2bee3e72b7e94f7f9a7568c4d7))
* **control:** Disable sha pin by default and update readme ([1d24837](https://github.com/stephrobert/plumber/commit/1d2483747880d2f6f53b87b64064975bffb313e1))
* **controls:** Fix bug in controls parsing and swap around some functions and files ([cdf0507](https://github.com/stephrobert/plumber/commit/cdf050792ce83b674ce80429f63767d81c90f321))
* **detection:** support SSH URL and Git protocol formats in remote auto-detection ([8e162aa](https://github.com/stephrobert/plumber/commit/8e162aaf7a9f21cedd183edf2a3aa00c6dc91b5d)), closes [#36](https://github.com/stephrobert/plumber/issues/36)
* **doc:** Add plumber to trusted images ([2a80e1a](https://github.com/stephrobert/plumber/commit/2a80e1a831d42b31d75a5af5407a2a2e7582473a))
* **license:** Update to use Elv2 license ([01656d0](https://github.com/stephrobert/plumber/commit/01656d0664524323264bd5cbd7d1cb3419e1f7ce))
* **naming:** Fix further naming convention with plumber ([3389f25](https://github.com/stephrobert/plumber/commit/3389f2581be70a079990490e0880b3f15160c972))
* **naming:** Rename to plumber and disable majors ([f442113](https://github.com/stephrobert/plumber/commit/f442113514cebd654431a80e2d30b2ea1289dbfa))
* **rebase:** Rebase on main and add spinner ([20e8c63](https://github.com/stephrobert/plumber/commit/20e8c63c5334875e467244fb55bcb059e978c3b8))
* **release:** empty commit to trigger release and push ([e8bd954](https://github.com/stephrobert/plumber/commit/e8bd954354e6c11ceef9fc53c89d634196a0e7af))
* **release:** Persist creds throughout release cycle ([4483cf0](https://github.com/stephrobert/plumber/commit/4483cf0fc8704f0763f95ee8bfc28462be66e157))
* **variables:** Fix self referential variable ([d5aa9a9](https://github.com/stephrobert/plumber/commit/d5aa9a93891b61c7a70a63f533d676084820a03a))


### ♻️ Refactoring

* **cmd:** extract config validation logic ([db02f52](https://github.com/stephrobert/plumber/commit/db02f524a28a83645631a0c7325bea161f1d86db))

## [0.1.52](https://github.com/getplumber/plumber/compare/v0.1.51...v0.1.52) (2026-03-04)


### ✨ Features

* **controls:** Add unsafe variable expansion control for user-filled predefined variables ([0f7ec72](https://github.com/getplumber/plumber/commit/0f7ec723ff0b03ffb1c345fa393eaa8f9e022dd7))

## [0.1.51](https://github.com/getplumber/plumber/compare/v0.1.50...v0.1.51) (2026-03-03)


### ✨ Features

* **controls:** Add debug trace detection control ([5e08b97](https://github.com/getplumber/plumber/commit/5e08b974d17ee2468768f6b8a2917b09eddfe0ff)), closes [#86](https://github.com/getplumber/plumber/issues/86)


### 🐛 Bug Fixes

* **rebase:** Rebase on main and add spinner ([20e8c63](https://github.com/getplumber/plumber/commit/20e8c63c5334875e467244fb55bcb059e978c3b8))

## [0.1.50](https://github.com/getplumber/plumber/compare/v0.1.49...v0.1.50) (2026-03-03)


### ✨ Features

* **cmd:** Add progress spinner during analysis ([80b1bad](https://github.com/getplumber/plumber/commit/80b1bad2a9b556d0e8414e2c90b4199cda5b26c4)), closes [#64](https://github.com/getplumber/plumber/issues/64)

## [0.1.49](https://github.com/getplumber/plumber/compare/v0.1.48...v0.1.49) (2026-03-03)


### ✨ Features

* **ci:** Replace trivvy with grype ([984286f](https://github.com/getplumber/plumber/commit/984286f2ca871792b8a12e34bea7a647cb72450f))

## [0.1.48](https://github.com/getplumber/plumber/compare/v0.1.47...v0.1.48) (2026-03-03)


### ✨ Features

* **ci:** Add dependabot ([3480a77](https://github.com/getplumber/plumber/commit/3480a777b3ef3bcf1dd72269da67aad28c3fcf0b))
* **ci:** Start adding test, lint, scan and pin versions by digest ([ec77e70](https://github.com/getplumber/plumber/commit/ec77e709864ddf621ca5337de3345c5c18627466))


### 🐛 Bug Fixes

* **ci:** Fix CI lint issues ([5288379](https://github.com/getplumber/plumber/commit/528837904a46b7b6334bbfe5f767d200faebdc81))

## [0.1.47](https://github.com/getplumber/plumber/compare/v0.1.46...v0.1.47) (2026-02-25)


### ✨ Features

* **controls:** Add overridden component and templates issue & integration into pbom and cyclonedex ([e970c27](https://github.com/getplumber/plumber/commit/e970c27697c29873057a5f6b6109715ba47d3896))

## [0.1.46](https://github.com/getplumber/plumber/compare/v0.1.45...v0.1.46) (2026-02-25)


### ✨ Features

* **cmd:** Add --fail-warnings on the analyze and config validate commands ([fbc5839](https://github.com/getplumber/plumber/commit/fbc58391ebc1e51d412588ed5748587a6e125227))
* **cmd:** validating config file before analysis ([a1dc4bd](https://github.com/getplumber/plumber/commit/a1dc4bd7c14d4fe230f770849c0b12683fba8323))


### ♻️ Refactoring

* **cmd:** extract config validation logic ([db02f52](https://github.com/getplumber/plumber/commit/db02f524a28a83645631a0c7325bea161f1d86db))

## [0.1.45](https://github.com/getplumber/plumber/compare/v0.1.44...v0.1.45) (2026-02-25)


### ✨ Features

* **cmd:** notify user when a newer version of plumber is available ([deca33f](https://github.com/getplumber/plumber/commit/deca33f74d9c2f1c2468d8ce98a5693df1fa9d9c)), closes [#39](https://github.com/getplumber/plumber/issues/39)
* **version:** async update check with opt-out ([a1ef745](https://github.com/getplumber/plumber/commit/a1ef74553e85c2341e803ea4ec683f6f0c4d6e31))


### 🐛 Bug Fixes

* **release:** Persist creds throughout release cycle ([4483cf0](https://github.com/getplumber/plumber/commit/4483cf0fc8704f0763f95ee8bfc28462be66e157))

## [0.1.45](https://github.com/getplumber/plumber/compare/v0.1.44...v0.1.45) (2026-02-25)


### ✨ Features

* **cmd:** notify user when a newer version of plumber is available ([deca33f](https://github.com/getplumber/plumber/commit/deca33f74d9c2f1c2468d8ce98a5693df1fa9d9c)), closes [#39](https://github.com/getplumber/plumber/issues/39)
* **version:** async update check with opt-out ([a1ef745](https://github.com/getplumber/plumber/commit/a1ef74553e85c2341e803ea4ec683f6f0c4d6e31))

## [0.1.44](https://github.com/getplumber/plumber/compare/v0.1.43...v0.1.44) (2026-02-19)


### ✨ Features

* **analyze:** add --controls and --skip-controls control filtering ([9a9aca0](https://github.com/getplumber/plumber/commit/9a9aca0a50e63eb27a7f7c8f5470ae6922ff425c))


### 🐛 Bug Fixes

* **controls:** Fix bug in controls parsing and swap around some functions and files ([cdf0507](https://github.com/getplumber/plumber/commit/cdf050792ce83b674ce80429f63767d81c90f321))

## [0.1.43](https://github.com/getplumber/plumber/compare/v0.1.42...v0.1.43) (2026-02-18)


### ✨ Features

* **config:** Add ValidateKnownKeys to warn on unknown config keys ([4c33ca3](https://github.com/getplumber/plumber/commit/4c33ca3c980cc57419b7b457bfae09c7543ba8d3)), closes [#58](https://github.com/getplumber/plumber/issues/58) [#58](https://github.com/getplumber/plumber/issues/58)


### 🐛 Bug Fixes

* **config:** Fix compilation issues + make validation recursive to test subkeys ([405abe4](https://github.com/getplumber/plumber/commit/405abe4dc7f78f2bee3e72b7e94f7f9a7568c4d7))

## [0.1.42](https://github.com/getplumber/plumber/compare/v0.1.41...v0.1.42) (2026-02-17)


### ✨ Features

* **analysis:** Add --mr-comment to create mr comments during analysis. Add --badge to create/update Plumber compliance badge when running on default remote branch ([4cba483](https://github.com/getplumber/plumber/commit/4cba4839ce3512ef2da7ff5b9512dca701d2a7f7))

## [0.1.41](https://github.com/getplumber/plumber/compare/v0.1.40...v0.1.41) (2026-02-12)


### ✨ Features

* **local:** Enable lint, validation and analysis of local .gitlab-ci.yml as well as local reoslution of include: local types. ([5a2a3aa](https://github.com/getplumber/plumber/commit/5a2a3aa3b67c8489289bf98f47f9494f386f6458))

## [0.1.40](https://github.com/getplumber/plumber/compare/v0.1.39...v0.1.40) (2026-02-12)


### ✨ Features

* **UX:** Integrate the control pinned by digest inside the immutable one ([87bd450](https://github.com/getplumber/plumber/commit/87bd45074c7d0710f67dfe23e486421eda1baf39))

## [0.1.39](https://github.com/getplumber/plumber/compare/v0.1.38...v0.1.39) (2026-02-12)


### ✨ Features

* **UX:** If a control is misisng from .plumber.yaml simply skip it instead of returning an error ([3eec388](https://github.com/getplumber/plumber/commit/3eec388d75d0e28792c580dc95c79f89663af5e3))

## [0.1.38](https://github.com/getplumber/plumber/compare/v0.1.37...v0.1.38) (2026-02-12)


### ✨ Features

* **controls:** add image digest pinning control ([ea538a9](https://github.com/getplumber/plumber/commit/ea538a954def56d4695b0d766f3bb1aff8ee7bbd))


### 🐛 Bug Fixes

* **control:** Disable sha pin by default and update readme ([1d24837](https://github.com/getplumber/plumber/commit/1d2483747880d2f6f53b87b64064975bffb313e1))

## [0.1.37](https://github.com/getplumber/plumber/compare/v0.1.36...v0.1.37) (2026-02-11)


### ✨ Features

* **artifact:** Add new concept: Pipeline Bill Of Materials (PBOM) and add cyclonedx output format support ([7097605](https://github.com/getplumber/plumber/commit/7097605f85123ea7599e67d7b22e34bfa13e726b))

## [0.1.36](https://github.com/getplumber/plumber/compare/v0.1.35...v0.1.36) (2026-02-11)


### ✨ Features

* **conf:** Add reference to examples in test file for required includes ([a8ec829](https://github.com/getplumber/plumber/commit/a8ec82961ab5152a3b6a1bff938afbc63978b7e6))

## [0.1.35](https://github.com/getplumber/plumber/compare/v0.1.34...v0.1.35) (2026-02-11)


### 🐛 Bug Fixes

* **detection:** support SSH URL and Git protocol formats in remote auto-detection ([8e162aa](https://github.com/getplumber/plumber/commit/8e162aaf7a9f21cedd183edf2a3aa00c6dc91b5d)), closes [#36](https://github.com/getplumber/plumber/issues/36)

## [0.1.34](https://github.com/getplumber/plumber/compare/v0.1.33...v0.1.34) (2026-02-10)


### ✨ Features

* **control:** Support Natural Language in pipeline inclusion for templates and components ([59c4edd](https://github.com/getplumber/plumber/commit/59c4edddb2750296d06d282e7c5efd272e3aa81d))

## [0.1.33](https://github.com/getplumber/plumber/compare/v0.1.32...v0.1.33) (2026-02-10)


### 🐛 Bug Fixes

* **branch:** use correct SHA for ciConfig query when --branch is specified ([1729084](https://github.com/getplumber/plumber/commit/1729084509de141d282e3a3d49c62fe7864e385e))

## [0.1.32](https://github.com/getplumber/plumber/compare/v0.1.31...v0.1.32) (2026-02-04)


### ✨ Features

* **control:** Make component collecetion compatible with gitlab built-in components ([532f071](https://github.com/getplumber/plumber/commit/532f071544c58f8c3af1cbf4771b43a1e296a799))

## [0.1.31](https://github.com/getplumber/plumber/compare/v0.1.30...v0.1.31) (2026-02-04)


### ✨ Features

* **control:** Add 3 new controls ([591a850](https://github.com/getplumber/plumber/commit/591a8509f47c1cc21eb3bf71ad185163368ba033))

## [0.1.30](https://github.com/getplumber/plumber/compare/v0.1.29...v0.1.30) (2026-02-03)


### ✨ Features

* **analysis:** Allow auto-detection for gitlab url and project during analysis + update banner ([e7a20e6](https://github.com/getplumber/plumber/commit/e7a20e6e2b49bdf7cc7805effebc765e73930b2d))

## [0.1.29](https://github.com/getplumber/plumber/compare/v0.1.28...v0.1.29) (2026-02-03)


### ✨ Features

* **conf:** Add conf view and move generate under conf ([8e549e9](https://github.com/getplumber/plumber/commit/8e549e97ab462dbf46d7f5d25dea8fd77989c796))

## [0.1.28](https://github.com/getplumber/plumber/compare/v0.1.27...v0.1.28) (2026-02-02)


### ✨ Features

* **update:** Empty commit ([b7bd04f](https://github.com/getplumber/plumber/commit/b7bd04fa8ab8430c15afd625ea27ce3ae2030e8e))

## [0.1.27](https://github.com/getplumber/plumber/compare/v0.1.26...v0.1.27) (2026-01-30)


### ✨ Features

* **ci:** Run on ubuntu 24.04 instead of latest ([b5473d2](https://github.com/getplumber/plumber/commit/b5473d2d527dc6672f332670d729f2eb4701944f))

## [0.1.26](https://github.com/getplumber/plumber/compare/v0.1.25...v0.1.26) (2026-01-30)


### ✨ Features

* **brew:** Test release 0.1.26 ([1848d52](https://github.com/getplumber/plumber/commit/1848d5248fd9c38cc54d8da11b24ef9694afb48d))

## [0.1.25](https://github.com/getplumber/plumber/compare/v0.1.24...v0.1.25) (2026-01-30)


### 🐛 Bug Fixes

* **brew:** Typo in release ([6c80575](https://github.com/getplumber/plumber/commit/6c80575c26f7294e8fa3a1553d10aa3f88864adb))
* **brew:** Typo in release ([4d7b905](https://github.com/getplumber/plumber/commit/4d7b905d84e47f25cda4c7b770b0f8660f453f7e))

## [0.1.24](https://github.com/getplumber/plumber/compare/v0.1.23...v0.1.24) (2026-01-30)


### ✨ Features

* **brew:** Enable automatic updating of brew tap formula repo upon new release ([ead9860](https://github.com/getplumber/plumber/commit/ead98601a2e355403975b087a5f0da1b3653fb76))

## [0.1.23](https://github.com/getplumber/plumber/compare/v0.1.22...v0.1.23) (2026-01-30)


### ✨ Features

* **conf:** Correct dockerfile and release file ([34263ec](https://github.com/getplumber/plumber/commit/34263ec71a90183a65e9b1511e90f34ce6d2e1a2))

## [0.1.22](https://github.com/getplumber/plumber/compare/v0.1.21...v0.1.22) (2026-01-30)


### ✨ Features

* **conf:** Allow conf generation with command ([7390e76](https://github.com/getplumber/plumber/commit/7390e763055c23a44506ba43db2331fb09a0b0e7))

## [0.1.21](https://github.com/getplumber/plumber/compare/v0.1.20...v0.1.21) (2026-01-29)


### ✨ Features

* **analyze:** Make conf and threshold optional ([bf6a4df](https://github.com/getplumber/plumber/commit/bf6a4dfc3c306d0bb9bd3cd8e7e6a20bcaf9522a))

## [0.1.20](https://github.com/getplumber/plumber/compare/v0.1.19...v0.1.20) (2026-01-28)


### ✨ Features

* **license:** Update license in readme to MPL-2.0 ([4cbab86](https://github.com/getplumber/plumber/commit/4cbab86370f2ee3c5d010d5888e17e1d6fa9d445))

## [0.1.19](https://github.com/getplumber/plumber/compare/v0.1.18...v0.1.19) (2026-01-23)


### 🐛 Bug Fixes

* **bug:** Cleanup some dead code ([fa7e1ae](https://github.com/getplumber/plumber/commit/fa7e1ae3f445ed2047c5ab4ef0a4a66f0bc8ad93))

## [0.1.18](https://github.com/getplumber/plumber/compare/v0.1.17...v0.1.18) (2026-01-23)


### ✨ Features

* **conf:** Introduce priority and automatic detection of conf files ([91ef31b](https://github.com/getplumber/plumber/commit/91ef31b3a43b1fa150a8fba1c2c880a2220b80ef))

## [0.1.17](https://github.com/getplumber/plumber/compare/v0.1.16...v0.1.17) (2026-01-23)


### ✨ Features

* **analysis:** Revert CI_JOB_TOKEN ([6c12fb5](https://github.com/getplumber/plumber/commit/6c12fb59ad73147ca96c3e58ce570eab751706eb))

## [0.1.16](https://github.com/getplumber/plumber/compare/v0.1.15...v0.1.16) (2026-01-23)


### 🐛 Bug Fixes

* **analysis:** If no controls ran (e.g., data collection failed), compliance is 0% - we can't verify anything ([7ec0e72](https://github.com/getplumber/plumber/commit/7ec0e72ca459539c4a8e4d4fdbc04faf01ddfae8))

## [0.1.15](https://github.com/getplumber/plumber/compare/v0.1.14...v0.1.15) (2026-01-23)


### ✨ Features

* **component:** Allow verbosity in component ([b59015a](https://github.com/getplumber/plumber/commit/b59015a9bd64213ca0b3cf21ea02b16e833eee13))

## [0.1.14](https://github.com/getplumber/plumber/compare/v0.1.13...v0.1.14) (2026-01-23)


### ✨ Features

* **controls:** Rename control outputs and config to make them more human-readable & Start using CI_JOB_TOKEN if in the CI ([6669707](https://github.com/getplumber/plumber/commit/66697073a13a337230f88a5ba1213def645f474d))

## [0.1.13](https://github.com/getplumber/plumber/compare/v0.1.12...v0.1.13) (2026-01-22)


### ✨ Features

* **log:** Improve logging experience ([426bcf8](https://github.com/getplumber/plumber/commit/426bcf817bc62c382b5487443832b49372cbe890))

## [0.1.12](https://github.com/getplumber/plumber/compare/v0.1.11...v0.1.12) (2026-01-22)


### ✨ Features

* **UX:** Define default output file, add output json example ([3dbfa1c](https://github.com/getplumber/plumber/commit/3dbfa1c9172b561e35ce52d834e794bc2fb08091))

## [0.1.11](https://github.com/getplumber/plumber/compare/v0.1.10...v0.1.11) (2026-01-22)


### ✨ Features

* **naming:** Rename components to plumber, no need for the analyze suffix ([53a0816](https://github.com/getplumber/plumber/commit/53a08165e7d879099606505c98cce7abc45715eb))

## [0.1.10](https://github.com/getplumber/plumber/compare/v0.1.9...v0.1.10) (2026-01-22)


### ✨ Features

* **output:** Improve readability of printed results ([97d708f](https://github.com/getplumber/plumber/commit/97d708f75d5b76fba6bb3724543837eb371f5c04))

## [0.1.9](https://github.com/getplumber/plumber/compare/v0.1.8...v0.1.9) (2026-01-21)


### 🐛 Bug Fixes

* **build:** Move release creation to after asset upload ([bc96e39](https://github.com/getplumber/plumber/commit/bc96e39f0a4fc56185b62199a1f0da8e6a503367))

## [0.1.8](https://github.com/getplumber/plumber/compare/v0.1.7...v0.1.8) (2026-01-21)


### ✨ Features

* **build:** Add platforms binary releases ([01d9bfa](https://github.com/getplumber/plumber/commit/01d9bfa79b80cf9d2c7ea6d1984e2acc7a4db9bb))

## [0.1.7](https://github.com/getplumber/plumber/compare/v0.1.6...v0.1.7) (2026-01-19)


### 🐛 Bug Fixes

* **analysis:** Fix bug where analyzed branch was being mistaken for branches to protect ([afdd5f8](https://github.com/getplumber/plumber/commit/afdd5f8424a3a05120e50feb3dd99909bd5dba6a))

## [0.1.6](https://github.com/getplumber/plumber/compare/v0.1.5...v0.1.6) (2026-01-19)


### 🐛 Bug Fixes

* **comment:** Add timeout comment to client ([15df3f0](https://github.com/getplumber/plumber/commit/15df3f002dea5c9ba5a6a95b9943eafa1dc0230d))

## [0.1.5](https://github.com/getplumber/plumber/compare/v0.1.4...v0.1.5) (2026-01-19)


### 🐛 Bug Fixes

* **component:** Add full docker path to plumber as trusted ([d7732c8](https://github.com/getplumber/plumber/commit/d7732c8d1fefd7f12730e60e963a703d2a120a64))

## [0.1.4](https://github.com/getplumber/plumber/compare/v0.1.3...v0.1.4) (2026-01-19)


### 🐛 Bug Fixes

* **doc:** Add plumber to trusted images ([2a80e1a](https://github.com/getplumber/plumber/commit/2a80e1a831d42b31d75a5af5407a2a2e7582473a))

## [0.1.3](https://github.com/getplumber/plumber/compare/v0.1.2...v0.1.3) (2026-01-19)


### 🐛 Bug Fixes

* **variables:** Fix self referential variable ([d5aa9a9](https://github.com/getplumber/plumber/commit/d5aa9a93891b61c7a70a63f533d676084820a03a))

## [0.1.2](https://github.com/getplumber/plumber/compare/v0.1.1...v0.1.2) (2026-01-19)


### ✨ Features

* **build:** Move to alpine to make command customizable in CI ([763bcf3](https://github.com/getplumber/plumber/commit/763bcf3eadd21fdf53503033b00ced44b1a6b862))
* **release:** Downgrade feat to patch ([eb30e81](https://github.com/getplumber/plumber/commit/eb30e8183466068954edbd4e700986e02bfd72af))

## [0.2.0](https://github.com/getplumber/plumber/compare/v0.1.1...v0.2.0) (2026-01-19)


### ✨ Features

* **build:** Move to alpine to make command customizable in CI ([763bcf3](https://github.com/getplumber/plumber/commit/763bcf3eadd21fdf53503033b00ced44b1a6b862))

## [0.1.1](https://github.com/getplumber/plumber/compare/v0.1.0...v0.1.1) (2026-01-19)


### 🐛 Bug Fixes

* **release:** empty commit to trigger release and push ([e8bd954](https://github.com/getplumber/plumber/commit/e8bd954354e6c11ceef9fc53c89d634196a0e7af))

## [0.0.1](https://github.com/getplumber/plumber/compare/v0.0.0...v0.0.1) (2026-01-19)


### 🐛 Bug Fixes

* **license:** Update to use Elv2 license ([01656d0](https://github.com/getplumber/plumber/commit/01656d0664524323264bd5cbd7d1cb3419e1f7ce))
* **naming:** Fix further naming convention with plumber ([3389f25](https://github.com/getplumber/plumber/commit/3389f2581be70a079990490e0880b3f15160c972))
* **naming:** Rename to plumber and disable majors ([f442113](https://github.com/getplumber/plumber/commit/f442113514cebd654431a80e2d30b2ea1289dbfa))
