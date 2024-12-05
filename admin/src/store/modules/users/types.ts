// src/store/modules/users/types.ts
declare namespace Api {
  namespace User {
    // Type for a single user
    export interface User {
      id: number;
      first_name: string;
      last_name: string;
      email: string;
      mobile: string;
      user_type: string;
    }

    // Type for a list of users
    // type UserList = User[];
    export type UserList = User[];

    interface FetchUsersResponse {
      status: number;
      message: string;
      data: User[];
    }

    // Payload for creating a new user
    export interface CreateUserPayload {
      first_name: string;
      last_name: string;
      email: string;
      mobile: string;
      user_type: string;
    }

    // Payload for updating an existing user
    export interface UpdateUserPayload {
      first_name?: string;
      last_name?: string;
      email?: string;
      mobile?: string;
      user_type?: string;
    }
    interface FetchUsersResponse {
      status: number;
      message: string;
      data: User[];
    }

    interface ApiResponse {
      response: {
        data: FetchUsersResponse;
      };
    }
  }

}
