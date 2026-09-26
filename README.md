# github-follower-prune

Cleans out mass-followers: accounts that follow you but themselves follow a huge number of people. It finds the followers of `USER_NAME` whose follower and following counts both exceed the configured thresholds, blocks them, then unblocks them so they stop following you without being left blocked.

## requirements

- Go 1.26+
- a GitHub token with permission to block users (see [permissions](#permissions))

## setup

1. `make init` creates `.env` from `.env.example`.
2. Put your token in `.env`:

   ```
   GH_TOKEN=...
   ```

3. Edit `config.json`:

   | key | meaning |
   | --- | --- |
   | `USER_NAME` | the account whose followers get pruned |
   | `FOLLOWERS_THRESHOLD_TO_BLOCK` | a follower is pruned when its own follower count is greater than this |
   | `FOLLOWING_THRESHOLD_TO_BLOCK` | and its own following count is greater than this |

   Both conditions must hold, so the defaults (`1000`/`1000`) only catch accounts that are large on both sides.

## usage

```
make run
```

There is no dry-run: matching followers are blocked immediately. Each pruned account is printed, per-account failures go to stderr and the run continues, and the last line reports the totals.

## permissions

The token needs write access to "Block another user" and read access to users/followers:

- fine-grained PAT: "Block another user" (write), plus read access to followers and profile
- classic PAT or OAuth token: the `user` scope
- or reuse `gh`: `gh auth refresh -h github.com -s user`, then `gh auth token`

The default `gh auth login` scopes do not include `user`, so blocking returns `403` until it is added.

## how it works

1. list every follower of `USER_NAME`, paginated at 100 per page
2. fetch each follower's profile
3. when both their `followers` and `following` counts exceed the thresholds, send `PUT` then `DELETE` to `/user/blocks/{login}`

Blocking removes the follow relationship; unblocking afterwards avoids leaving the account blocked. Profile lookups run with a small concurrency cap to stay clear of secondary rate limits.

## development

```
make lint     # golangci-lint
make format   # goimports + golangci-lint fmt
```
