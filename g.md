Usage: swagger \[OPTIONS\] generate cli \[cli-OPTIONS\]

generate a command line client tool from the swagger spec

Application Options: -q, –quiet silence logs –log-output=LOG-FILE
redirect logs to file

Help Options: -h, –help Show this help message

\[cli command options\] -c, –client-package= the package to save the
client specific code (default: client) -P, –principal= the model to use
for the security principal –default-scheme= the default scheme for this
API (default: http) –principal-is-interface the security principal
provided is an interface-

                                                                                  , not a
                                                                                  struct
          --default-produces=                                                     the
                                                                                  default
                                                                                  mime type
                                                                                  that API
                                                                                  operation-

                                                                                  s produce
                                                                                  (default:
                                                                                  applicati-

                                                                                  on/json)
          --default-consumes=                                                     the
                                                                                  default
                                                                                  mime type
                                                                                  that API
                                                                                  operation-

                                                                                  s consume
                                                                                  (default:
                                                                                  applicati-

                                                                                  on/json)
          --skip-models                                                           no models
                                                                                  will be
                                                                                  generated
                                                                                  when this
                                                                                  flag is
                                                                                  specified
          --skip-operations                                                       no
                                                                                  operation-

                                                                                  s will be
                                                                                  generated
                                                                                  when this
                                                                                  flag is
                                                                                  specified
      -A, --name=                                                                 the name
                                                                                  of the
                                                                                  applicati-

                                                                                  on,
                                                                                  defaults
                                                                                  to a
                                                                                  mangled
                                                                                  value of
                                                                                  info.title
          --cli-app-name=                                                         the app
                                                                                  name for
                                                                                  the cli
                                                                                  executabl-

                                                                                  e. useful
                                                                                  for go
                                                                                  install.
                                                                                  (default:
                                                                                  cli)

    Options common to all code generation commands:
      -f, --spec=                                                                 the spec
                                                                                  file to
                                                                                  use
                                                                                  (default
                                                                                  swagger.{-

                                                                                  json,yml,-

                                                                                  yaml})
      -t, --target=                                                               the base
                                                                                  directory
                                                                                  for
                                                                                  generatin-

                                                                                  g the
                                                                                  files
                                                                                  (default:
                                                                                  ./)
          --template=[stratoscale]                                                load
                                                                                  contribut-

                                                                                  ed
                                                                                  templates
      -T, --template-dir=                                                         alternati-

                                                                                  ve
                                                                                  template
                                                                                  override
                                                                                  directory
      -C, --config-file=                                                          configura-

                                                                                  tion file
                                                                                  to use
                                                                                  for
                                                                                  overridin-

                                                                                  g
                                                                                  template
                                                                                  options
      -r, --copyright-file=                                                       copyright
                                                                                  file used
                                                                                  to add
                                                                                  copyright
                                                                                  header
          --additional-initialism=                                                consecuti-

                                                                                  ve
                                                                                  capitals
                                                                                  that
                                                                                  should be
                                                                                  considere-

                                                                                  d
                                                                                  intialisms
          --allow-template-override                                               allows
                                                                                  overridin-

                                                                                  g
                                                                                  protected
                                                                                  templates
          --skip-validation                                                       skips
                                                                                  validatio-

                                                                                  n of spec
                                                                                  prior to
                                                                                  generation
          --dump-data                                                             when
                                                                                  present
                                                                                  dumps the
                                                                                  json for
                                                                                  the
                                                                                  template
                                                                                  generator
                                                                                  instead
                                                                                  of
                                                                                  generatin-

                                                                                  g files
          --strict-responders                                                     Use
                                                                                  strict
                                                                                  type for
                                                                                  the
                                                                                  handler
                                                                                  return
                                                                                  value
          --with-expand                                                           expands
                                                                                  all
                                                                                  $ref's in
                                                                                  spec
                                                                                  prior to
                                                                                  generatio-

                                                                                  n
                                                                                  (shorthan-

                                                                                  d to
                                                                                  --with-fl-

                                                                                  atten=exp-

                                                                                  and)
          --with-flatten=[minimal|full|expand|verbose|noverbose|remove-unused]    flattens
                                                                                  all
                                                                                  $ref's in
                                                                                  spec
                                                                                  prior to
                                                                                  generatio-

                                                                                  n
                                                                                  (default:
                                                                                  minimal,
                                                                                  verbose)
      -p, --template-plugin=                                                      the
                                                                                  template
                                                                                  plugin to
                                                                                  use

    Options for model generation:
      -m, --model-package=                                                        the
                                                                                  package
                                                                                  to save
                                                                                  the
                                                                                  models
                                                                                  (default:
                                                                                  models)
      -M, --model=                                                                specify a
                                                                                  model to
                                                                                  include
                                                                                  in
                                                                                  generatio-

                                                                                  n, repeat
                                                                                  for
                                                                                  multiple
                                                                                  (defaults
                                                                                  to all)
          --existing-models=                                                      use
                                                                                  pre-gener-

                                                                                  ated
                                                                                  models
                                                                                  e.g.
                                                                                  github.co-

                                                                                  m/foobar/-

                                                                                  model
          --strict-additional-properties                                          disallow
                                                                                  extra
                                                                                  propertie-

                                                                                  s when
                                                                                  additiona-

                                                                                  lProperti-

                                                                                  es is set
                                                                                  to false
          --keep-spec-order                                                       keep
                                                                                  schema
                                                                                  propertie-

                                                                                  s order
                                                                                  identical
                                                                                  to spec
                                                                                  file
          --struct-tags=                                                          the
                                                                                  struct
                                                                                  tags to
                                                                                  generate,
                                                                                  repeat
                                                                                  for
                                                                                  multiple
                                                                                  (defaults
                                                                                  to json)

    Options for operation generation:
      -O, --operation=                                                            specify
                                                                                  an
                                                                                  operation
                                                                                  to
                                                                                  include,
                                                                                  repeat
                                                                                  for
                                                                                  multiple
                                                                                  (defaults
                                                                                  to all)
          --tags=                                                                 the tags
                                                                                  to
                                                                                  include,
                                                                                  if not
                                                                                  specified
                                                                                  defaults
                                                                                  to all
      -a, --api-package=                                                          the
                                                                                  package
                                                                                  to save
                                                                                  the
                                                                                  operation-

                                                                                  s
                                                                                  (default:
                                                                                  operation-

                                                                                  s)
          --with-enum-ci                                                          allow
                                                                                  case-inse-

                                                                                  nsitive
                                                                                  enumerati-

                                                                                  ons
          --skip-tag-packages                                                     skips the
                                                                                  generatio-

                                                                                  n of
                                                                                  tag-based
                                                                                  operation
                                                                                  packages,
                                                                                  resulting
                                                                                  in a flat
                                                                                  generation
