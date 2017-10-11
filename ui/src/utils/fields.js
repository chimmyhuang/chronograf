import _ from 'lodash'
j
// functions returns all top-level fields with type
export const ofType = (fields, type) =>
  _.filter(fields, f => _.get(f, 'type') === type)

// functions returns all top-level functions in fields
export const functions = fields => ofType(fields, 'func')

// numFunctions searches queryConfig fields for functions
export const numFunctions = fields => _.size(functions(fields))

// functionNames returns the names of all top-level functions
export const functionNames = fields => _.map(functions(fields), f => f.name)

// firstFieldName returns the name of the first of type field
export const firstFieldName = fields => _.head(fieldNamesDeep(fields))

// getFields returns all of the top-level fields of type field
export const getFields = fields => ofType(fields, 'field')

export const fieldsDeep = fields =>
  _.uniqBy(walk(fields, f => getFields(f)), 'name')

export const fieldNamesDeep = fields =>
  _.map(fieldsDeep(fields), f => _.get(f, 'name'))

export const walk = (fields, fn) =>
  _.foreach(fields, f => _.concat(fn(f), walk(_.get(f, 'args', []), fn)))