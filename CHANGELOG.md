# Changelog

## 0.6.0 (2026-03-16)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/keycardai/keycard-go/compare/v0.5.0...v0.6.0)

### Features

* add OAuth2 client credentials security scheme from common spec ([7c9c2ff](https://github.com/keycardai/keycard-go/commit/7c9c2ffdd29dfa01475970e322ab07fca5922a8c))
* update pkg-oapi-common and add OAuth2 security scheme ([7be0447](https://github.com/keycardai/keycard-go/commit/7be0447b0ed9c43a102abf16cbaebce593d99e66))
* use common bearerAuth and OAuth2 security schemes ([598c4fd](https://github.com/keycardai/keycard-go/commit/598c4fdba1349d751883ee840a7b943c96c6c49f))


### Chores

* configure new SDK language ([baff3ed](https://github.com/keycardai/keycard-go/commit/baff3edd0210f28cfdcf49344c52d358eecdb5aa))
* **internal:** tweak CI branches ([8e63a5a](https://github.com/keycardai/keycard-go/commit/8e63a5a4b1ff3793c0575550bfe6cb984b11660b))


### Documentation

* remove MCP endpoints ([a391a65](https://github.com/keycardai/keycard-go/commit/a391a65b98177875d5f9ab336bc8a3f385a3fd06))

## 0.5.0 (2026-03-16)

Full Changelog: [v0.4.1...v0.5.0](https://github.com/keycardai/keycard-go/compare/v0.4.1...v0.5.0)

### Features

* add OAuth2 as alternative auth on management API endpoints ([55f93cc](https://github.com/keycardai/keycard-go/commit/55f93ccb8c595582d1728d508ebf714f2c679678))
* add OAuth2 client_credentials auth to SDK config ([146cd2d](https://github.com/keycardai/keycard-go/commit/146cd2df49ec79e76eecec89f35444bc41a3c98c))
* consolidate prefixed security schemes into canonical names ([5ea0b3a](https://github.com/keycardai/keycard-go/commit/5ea0b3a8c5e6fb323b85c5583e9e1acaacd06084))
* Include `array_format: brackets` settings ([b54c425](https://github.com/keycardai/keycard-go/commit/b54c425b7f7d68cbb2a2a8e31eb817675e203010))
* remove unused security schemes from joined spec ([a6b1b2e](https://github.com/keycardai/keycard-go/commit/a6b1b2e54d20c92a8d9a8c3a547c4920ab5d5480))
* support HTTP Basic Auth for service account token endpoint (RFC 6749 2.3.1) ([476e7ba](https://github.com/keycardai/keycard-go/commit/476e7bab7abccd75e8ab4b561f6e6f4585958196))


### Bug Fixes

* **tests:** prevent tests failing due to making unnecessary OAuth token requests ([286daf8](https://github.com/keycardai/keycard-go/commit/286daf841ec1794bf74dfee47bb199cfbb848c50))


### Chores

* hide unstable mcp features from api documentation ([361cffc](https://github.com/keycardai/keycard-go/commit/361cffc3611dbcf7acebc76b248babe534a966bc))
* **internal:** use explicit returns in more places ([03e2e8f](https://github.com/keycardai/keycard-go/commit/03e2e8fee7da79e8a30019dc29b17e745a5d9bf5))

## 0.4.1 (2026-03-10)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/keycardai/keycard-go/compare/v0.4.0...v0.4.1)

### Chores

* **internal:** use explicit returns ([9acde8f](https://github.com/keycardai/keycard-go/commit/9acde8f9cfbbe25b1620a79bb9362f99cd8b8836))

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
