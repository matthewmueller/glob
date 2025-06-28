# 0.1.1 / 2025-06-28

- fix panic due to yielding after range function exits

# 0.1.0 / 2025-06-23

- large breaking change
- separated into globfs and glob.
- renamed: match to glob
- added: is, match
- most apis now accept multiple patterns
- smarter file walking based on greatest common ancestor

# 0.0.5 / 2025-02-17

- fix expand to support multiple patterns

# 0.0.4 / 2024-08-22

- expose glob.Base
- support passing in file separators

# 0.0.3 / 2024-03-16

- incorporate multiple base dir support in walk

# 0.0.2 / 2024-02-19

- add glob.WalkFS and glob.MatchFS

# 0.0.1 / 2024-02-12

- add license, readme and example
- modernize the glob package

# Untagged / 2018-11-30

- watch setup but not implements
- Initial commit
