RSS Reader
==========

Fetch RSS feeds, log them to a file and show in a simple CGI interface.

Download and run
----------------

See the [Releases](https://github.com/janosgyerik/rssreader/releases) tab on GitHub for binaries.

Create a configuration like [this](https://github.com/janosgyerik/rssreader/blob/master/feeds.yml.example),
and save it in a file named `feeds.yml`:

    feeds:
      - id: so-java-bounty
        url: http://stackoverflow.com/feeds/tag?tagnames=java&sort=featured
      - id: so-go-bounty
        url: http://stackoverflow.com/feeds/tag?tagnames=go&sort=featured

Edit the list of `feeds`:

- The `id` can be arbitrary, it is only used for display purposes.
- The URL should be an RSS feed.

Develop
-------

Download dependencies:

    go get

Create configuration: copy `feeds.yml.example` to `feeds.yml` and customize.

Run using the default configuration file (`feeds.yml`):

    ./run.sh

To generate test coverage report, run:

    ./coverage.sh

To rebuild the binaries for multiple platforms, run:

    ./rebuild-all.sh
