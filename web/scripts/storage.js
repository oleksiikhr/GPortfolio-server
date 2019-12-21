'use strict'

import { isObject, wrap as objWrap } from '@/scripts/obj'
import { wrap as arrWrap } from '@/scripts/arr'

/* | ---------------------------------------------------------------------------------------
 * | - Reserved keys -
 * | ---------------------------------------------------------------------------------------
 */

export const GITHUB_WEBSITE = 'website:github'

/* | ---------------------------------------------------------------------------------------
 * | - Functions -
 * | ---------------------------------------------------------------------------------------
 */

/**
 * @param {string} key
 * @returns {Array}
 */
export function getArray(key) {
  const item = localStorage.getItem(key)

  if (item) {
    try {
      return arrWrap(JSON.parse(item))
    } catch (e) {
      return []
    }
  }

  return []
}

/**
 * @param {string} key
 * @param {*} value
 * @returns {void}
 */
export function setArray(key, value) {
  localStorage.setItem(key, JSON.stringify(arrWrap(value)))
}

/**
 * @param {string} key
 * @returns {string}
 */
export function getString(key) {
  const item = localStorage.getItem(key)

  if (!item || isObject(item)) {
    return ''
  }

  if (typeof item === 'number') {
    return item.toString()
  }

  return item || ''
}

/**
 * @param {string} key
 * @param {*} value
 */
export function setString(key, value) {
  localStorage.setItem(key, value)
}

/**
 * @param {string} key
 * @param value
 */
export function setObject(key, value) {
  if (value && isObject(value)) {
    localStorage.setItem(key, JSON.stringify(value))
  }
}

/**
 * @param {string} key
 * @returns {object}
 */
export function getObject(key) {
  const item = localStorage.getItem(key)

  if (item) {
    try {
      return objWrap(JSON.parse(item))
    } catch (e) {
      return {}
    }
  }

  return {}
}
