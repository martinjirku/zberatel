import type { NhostClient } from '@nhost/nhost-js';

export class UserAuth {
  private client: NhostClient;
  public id?: string;
  constructor(client: NhostClient) {
    this.client = client;
  }
  async login(email: string, password: string) {
    const login = await this.client.auth.signIn({
      email: email,
      password: password,
    });

    return login;
  }
}
