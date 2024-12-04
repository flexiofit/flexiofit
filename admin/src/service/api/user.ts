import { request } from '../request';

const USER_BASE = '/user';

/**
 * Fetch all users
 *
 * @returns Promise resolving to the list of users
 */
export function fetchAllUsers() {
  return request<Api.User.UserList[]>({
    url: USER_BASE,
    method: 'get',
  });
}

/**
 * Fetch user by ID
 *
 * @param id - The ID of the user to fetch
 * @returns Promise resolving to the user data
 */
export function fetchUserById(id: string) {
  return request<Api.User.User>({
    url: `${USER_BASE}/${id}`,
    method: 'get',
  });
}

/**
 * Create a new user
 *
 * @param user - The user data to create
 * @returns Promise resolving to the created user
 */
export function createUser(user: Api.User.CreateUserPayload) {
  return request<Api.User.User>({
    url: USER_BASE,
    method: 'post',
    data: user,
  });
}

/**
 * Update an existing user
 *
 * @param id - The ID of the user to update
 * @param user - The updated user data
 * @returns Promise resolving to the updated user
 */
export function updateUser(id: string, user: Api.User.UpdateUserPayload) {
  return request<Api.User.User>({
    url: `${USER_BASE}/${id}`,
    method: 'put',
    data: user,
  });
}

/**
 * Delete a user
 *
 * @param id - The ID of the user to delete
 * @returns Promise resolving to a success response
 */
export function deleteUser(id: string) {
  return request<void>({
    url: `${USER_BASE}/${id}`,
    method: 'delete',
  });
}
