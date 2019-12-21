'use strict'

/** @type {RegExp} */
const REGEX_PASS = /[0-9a-z]{70}@[0-9a-z]{30}/i

/**
 * Add an event to close the modal window and check the
 * data for login and password
 *
 * @param {Window} w
 * @returns {Promise<>}
 */
export function windowAuth(w) {
  return new Promise((resolve, reject) => {
    w.addEventListener('beforeunload', (evt) => {
      const response = evt.target.body.innerText

      if (REGEX_PASS.test(response)) {
        const [login, password] = response.split('@')
        // TODO Notify status
        resolve({ login, password })
      }

      // TODO Notify status
      reject()
    })
  })
}

/**
 * Get headers for auth request
 *
 * @param {string} login
 * @param {string} password
 * @returns {{'Security-Key': string, 'Security-Pass': string | string}}
 */
export function getAuthHeaders(login, password) {
  return {
    'Security-Pass': password,
    'Security-Key': login
  }
}

export default {
  windowAuth,
  getAuthHeaders
}
