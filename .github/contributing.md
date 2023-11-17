# Contributing to `gokaf`

Thanks for checking out the gokaf. We're excited that you decided to
contribute and make gokaf awesome!

Following guidelines can help you figure out where you can best be helpful.

## Table of Contents

- [Types of contributions we're looking for](#types-of-contributions-were-looking-for)
- [Code of conduct](#code-of-conduct)
- [How to contribute](#how-to-contribute)
- [Style guide](#style-guide)


## Types of contributions we're looking for
- Report issues or problems that you face while using gokaf
- Adding more [examples](../examples)
- Improving the documentation and developer experience for fellow developers

Interested in making a contribution? Read on!

## Code of conduct

Before we get started, please note that we expect you to follow the open
source code of conduct and would appreciate if you abide by them. We in return,
promise that we will abide with the same to the best of our capacity.

The code of conduct is adopted from
[Contributor Covenant](https://www.contributor-covenant.org/).

## How to contribute

If you'd like to contribute, start by searching through the
[issues](https://github.com/acjzz/gokaf/issues) and
[pull requests](https://github.com/acjzz/gokaf/pulls) to see
whether someone else has raised a similar idea or question.

If you don't see your idea listed, please create an issue and discuss with
us. We promise to respond back and help out. In case this is related to any
enhancement or addition of core functionality, please discuss the approach
before starting to code; we might be working on similar things or have
suggestions for you.

We follow the Github - fork and pull workflow.

1. [Fork the repository](https://help.github.com/en/articles/fork-a-repo) and
clone your fork repository - to be named as `origin`
1. Add the remote repository - `git remote add upstream git@github.com:acjzz/gokaf.git`
    - In case the repository was already forked, update the develop branch from
    the latest merge using
    `git checkout develop && git pull upstream develop && git push origin develop`
    following the conventions
1. Create a feature or fix branch - `git checkout -b feature/fooBar`
1. Commit your changes - `git commit -am 'Feature(fooBar): Add some foo bar'`
1. Push feature/fix branch to `origin` repository
1. [Create a new pull request](https://help.github.com/en/articles/creating-a-pull-request)
to merge `origin:feature/fooBar` (right side) on `upstream:develop` (left side)
while adhering to the checklist and guidelines

While the prerequisites above must be satisfied prior to having your pull
request reviewed, the reviewer(s) may ask you to complete additional design
work, tests, or other changes before your pull request can be ultimately
accepted.

## Style guide

Please format your golang code using `gofmt` tool.
