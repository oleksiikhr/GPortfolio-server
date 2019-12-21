'use strict'

/**
 * Get deep value from object
 * @param {*} data
 * @param {array} keys
 * @param {number} deep
 * @return {*} undefined - not found
 */
export function getDeepByKey(data, keys, deep = 0) {
  if (!keys.length || keys.length === deep) {
    return data
  }

  if (!isObject(data)) {
    return undefined
  }

  const val = data[keys[deep]]

  if (isObject(val)) {
    return getDeepByKey(val, keys, deep + 1)
  }

  if (keys.length !== deep + 1) {
    return undefined
  }

  return val
}

/**
 * Value is an object
 * @param {*} val
 * @return {boolean}
 */
export function isObject(val) {
  return !!val && val.constructor === Object
}

/**
 * Convert data to an object (if they are not)
 *
 * @param {any} data
 * @returns {array}
 */
export function wrap(data) {
  return isObject(data) ? data : [data]
}

export default {
  getDeepByKey,
  isObject,
  wrap
}
