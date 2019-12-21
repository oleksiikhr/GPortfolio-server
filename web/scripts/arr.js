'use strict'

/**
 * Shuffles array in place. ES6 version
 * @param {array} arr items An array containing the items
 * @see https://stackoverflow.com/a/6274381
 */
export function shuffle(arr) {
  for (let i = arr.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [arr[i], arr[j]] = [arr[j], arr[i]]
  }

  return arr
}

/**
 * Convert data to an array (if they are not)
 *
 * @param {any} data
 * @returns {array}
 */
export function wrap(data) {
  return Array.isArray(data) ? data : [data]
}

export default {
  shuffle,
  wrap
}
