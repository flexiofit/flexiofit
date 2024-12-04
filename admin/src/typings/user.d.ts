declare namespace Api {
  namespace User {
    // Type for a single user
    interface User {
      id: string;
      name: string;
      email: string;
      role: string;
      createdAt: string;
      updatedAt: string;
    }

    // Type for a list of users
    type UserList = User[];

    // Payload for creating a new user
    interface CreateUserPayload {
      name: string;
      email: string;
      password: string;
      role: string;
    }

    // Payload for updating an existing user
    interface UpdateUserPayload {
      name?: string;
      email?: string;
      role?: string;
    }
  }
}
