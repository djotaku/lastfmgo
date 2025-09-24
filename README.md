# lastfmgo

This is a library for accessing the [last.fm API](https://www.last.fm/api) via golang. Currently it implements the endpoints I need for [lastfmSocials](https://github.com/djotaku/lastfmSocials) which took over for the taskign in [lastfmBluesky]( https://github.com/djotaku/lastfmbluesky) and [lastfmMastoston](https://github.com/djotaku/lastfmmastodon). I may eventually cover the full API, but it's not a focus for me. I will definitely accept PRs to add other endpoints, but I'm not ready for that yet. I need to make a few changes to this code first.

## API Stability

As I'm slowly adding other endpoints, I may realize that I need to change things with the API in order to keep things DRY. Until I get to a 1.x, make use of go.mod's dependency pinning and only upgrade once you check what has changed with the API.