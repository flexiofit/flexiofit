// The store implementation
import { defineStore } from 'pinia';
import {
  fetchAllUsers,
  fetchUserById,
  createUser,
  updateUser,
  deleteUser
} from '@/service/api/user';

export interface UserState {
  users: Api.User.UserList;
  currentUser: Api.User.User | null;
  loading: boolean;
  error: string | null;
}


export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    users: [],
    currentUser: null,
    loading: false,
    error: null
  }),

  getters: {
    getUsers: (state): Api.User.UserList => state.users,
    getCurrentUser: (state): Api.User.User | null => state.currentUser,
    isLoading: (state): boolean => state.loading,
    getError: (state): string | null => state.error
  },

  actions: {
    // async fetchUsers() {
    //   try {
    //     this.loading = true;
    //     const response = await fetchAllUsers(); // Fetch the full response
    //     console.log("Fetched Response:", response); // Log full response for debugging
    //
    //     if (response?.response?.data?.data) {
    //       this.users = response.response.data.data; // Assign the user list from nested data
    //     } else {
    //       this.error = "No users found in response";
    //       console.warn("Unexpected Response Format:", response);
    //     }
    //   } catch (err) {
    //     this.error = err instanceof Error ? err.message : 'Failed to fetch users';
    //     console.error("Error fetching users:", err); // Log error for debugging
    //     throw err;
    //   } finally {
    //     this.loading = false;
    //   }
    // },
    async fetchUsers() {
      try {
        this.loading = true;
        const response = await fetchAllUsers();
        console.log("Fetched Response:", response);

        // Safely assign the response data to users
        const userList = response.response?.data?.data as Api.User.UserList | undefined;
        if (userList) {
          this.users = userList; // Ensure correct type assignment
        } else {
          this.error = "No users found in response";
          console.warn("Unexpected Response Format:", response);
        }
      } catch (err) {
        this.error = err instanceof Error ? err.message : 'Failed to fetch users';
        console.error("Error fetching users:", err); // Log error for debugging
        throw err;
      } finally {
        this.loading = false;
      }
    },
    async fetchUserById(id: string) {
      try {
        this.loading = true;
        const response = await fetchUserById(id);
        if (response.data) {
          this.currentUser = response.data;
        }
        return response.data;
      } catch (err) {
        this.error = err instanceof Error ? err.message : 'Failed to fetch user';
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async createNewUser(userData: Api.User.CreateUserPayload) {
      try {
        this.loading = true;
        const response = await createUser(userData);
        if (response.data) {
          this.users.push(response.data);
        }
        return response.data;
      } catch (err) {
        this.error = err instanceof Error ? err.message : 'Failed to create user';
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async updateExistingUser(id: string, userData: Api.User.UpdateUserPayload) {
      try {
        this.loading = true;
        const response = await updateUser(id, userData);
        if (response.data) {
          const index = this.users.findIndex(user => user.id === Number(id));
          if (index !== -1) {
            this.users[index] = response.data;
          }
          if (this.currentUser?.id === Number(id)) {
            this.currentUser = response.data;
          }
        }
        return response.data;
      } catch (err) {
        this.error = err instanceof Error ? err.message : 'Failed to update user';
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async removeUser(id: string) {
      try {
        this.loading = true;
        await deleteUser(id);
        this.users = this.users.filter(user => user.id !== Number(id));
        if (this.currentUser?.id === Number(id)) {
          this.currentUser = null;
        }
      } catch (err) {
        this.error = err instanceof Error ? err.message : 'Failed to delete user';
        throw err;
      } finally {
        this.loading = false;
      }
    },

    clearError() {
      this.error = null;
    }
  }
});
