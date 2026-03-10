# Changelog

## 0.4.0 (2026-03-10)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/keycardai/keycard-go/compare/v0.3.0...v0.4.0)

### Features

* Typescript package name @keycardai/api ([240e21f](https://github.com/keycardai/keycard-go/commit/240e21fb569deedb5af81d2a45ac42e15f9ee7c4))

## 0.3.0 (2026-03-10)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/keycardai/keycard-go/compare/v0.2.0...v0.3.0)

### Features

* jelmer/stainless keycardai configuration ([1b5920e](https://github.com/keycardai/keycard-go/commit/1b5920e3b038fa6217d18fcc69146b360a7c8e72))


### Chores

* sync repo ([f2c6b51](https://github.com/keycardai/keycard-go/commit/f2c6b519df73135eaa2e6eca4c602a652ad376aa))

## 0.2.0 (2026-03-07)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/keycardlabs/keycard-go/compare/v0.1.0...v0.2.0)

### Features

* Add application_type property to url credential ([2cc792d](https://github.com/keycardlabs/keycard-go/commit/2cc792d7977e1d93f1f76596bcc9b30d59379bb2))
* add is_system_managed property to resources and applications ([6c016c2](https://github.com/keycardlabs/keycard-go/commit/6c016c27cc23c410a916cee827dedecd896c695e))
* filter, search, sort for invitations (COR-577) ([fcbe7f3](https://github.com/keycardlabs/keycard-go/commit/fcbe7f3bf4dc79aedb1b9c929b628dd7d7093f03))
* internal endpoint to re-send verification email ([364d718](https://github.com/keycardlabs/keycard-go/commit/364d71847dc7eeebcb9cb1d53697c693772fdccf))
* switch zone setting for email/password support (COR-543) ([c3113c0](https://github.com/keycardlabs/keycard-go/commit/c3113c0356e2c6197507f08283a3b4db737a848a))


### Bug Fixes

* fix request delays for retrying to be more respectful of high requested delays ([1a67996](https://github.com/keycardlabs/keycard-go/commit/1a67996350c34debf13cf0e8efb465fb563d0f07))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([261b573](https://github.com/keycardlabs/keycard-go/commit/261b573fd74b4e6f79e5d1501006bfcb5b2edc58))
* **internal:** codegen related update ([a952d28](https://github.com/keycardlabs/keycard-go/commit/a952d28ab82ac826806f8adf4b2b4246b6c6efee))
* **internal:** move custom custom `json` tags to `api` ([1d87561](https://github.com/keycardlabs/keycard-go/commit/1d87561f4612e73104d1b15b27081a314d4becef))
* sync repo ([fbb71a8](https://github.com/keycardlabs/keycard-go/commit/fbb71a80a3230fc70594844365f91f5eeb78ecb1))

## 0.1.0 (2026-02-13)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/keycardlabs/keycard-go/compare/v0.0.1...v0.1.0)

### Features

* PREVIEW: sync svc-iam.yaml from svc-iam/jelmer/updated-openapi-sync-workflow ([7054ad3](https://github.com/keycardlabs/keycard-go/commit/7054ad3abda358d771af345aa9e676d18889acd6))
* PREVIEW: sync svc-vault-api.yaml from svc-vault-api/jelmer/cor-298-add-openapi-sync ([0370f5b](https://github.com/keycardlabs/keycard-go/commit/0370f5badb67988db2535f81f2cfdf4920730e87))


### Bug Fixes

* **encoder:** correctly serialize NullStruct ([d2a9e01](https://github.com/keycardlabs/keycard-go/commit/d2a9e0121c10bb47f5ca2dabffaa4ca87f78e002))
* remove refs to internal components when stripping x-internal ([ce79f53](https://github.com/keycardlabs/keycard-go/commit/ce79f53509b3af4b660e5624275929458f1becd6))


### Chores

* configure new SDK language ([c7d1769](https://github.com/keycardlabs/keycard-go/commit/c7d17699f1d5dfd652b6eb4a2cd4997df883dcb3))
* update SDK settings ([3f27d6c](https://github.com/keycardlabs/keycard-go/commit/3f27d6c35a06402f957fe7cf235a14ccb3097eb9))
