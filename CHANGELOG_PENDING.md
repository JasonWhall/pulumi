### Improvements

- [dotnet/sdk] Support microsoft logging extensions with inline programs
  [#7117](https://github.com/pulumi/pulumi/pull/7117)

- [dotnet/sdk] Add create unknown to output utilities.
  [#7173](https://github.com/pulumi/pulumi/pull/7173)

- [dotnet] Fix Resharper code issues.
  [#7178](https://github.com/pulumi/pulumi/pull/7178)

- [codegen] - Include properties with an underlying type of string on Go provider instances.

- [cli] - Provide a more helpful error instead of panicking when codegen fails during import.
  [#7265](https://github.com/pulumi/pulumi/pull/7265)

- [codegen/python] - Cache package version for improved performance.
  [#7293](https://github.com/pulumi/pulumi/pull/7293)

- [sdk/python] - Reduce `log.debug` calls for improved performance
  [#7295](https://github.com/pulumi/pulumi/pull/7295)

### Bug Fixes

- [sdk/python] - Fix regression in behaviour for `Output.from_input({})`

- [sdk/python] - Fix hanging deployments and improve error messages
  for programs with incorrect typings for output values
  [#7049](https://github.com/pulumi/pulumi/pull/7049)

- [sdk/python] - Prevent infinite loops when iterating `Output` objects
  [#7288](https://github.com/pulumi/pulumi/pull/7288)

- [codegen/python] - Rename conflicting ResourceArgs classes
  [#7171](https://github.com/pulumi/pulumi/pull/7171)
