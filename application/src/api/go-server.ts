import type { AxiosInstance } from 'axios';
import axios from 'axios';

class GoServer {
  client: AxiosInstance;

  constructor() {
    this.client = axios.create({
      baseURL: 'http://localhost:8080',
      timeout: 1000,
      headers: {
        'Content-Type': 'application/json',
      },
    });
  }

  setAuthToken(token: string) {
    localStorage.setItem('authToken', token);
  }

  getAuthToken() {
    return localStorage.getItem('authToken');
  }

  removeAuthToken() {
    localStorage.removeItem('authToken');
  }

  async logIn(username: string, password: string) {
    const body = {
      username: username,
      password: password,
    };

    const response = await this.client.post('/auth/login', body);
    return response;
  }

  async signUp(username: string, email: string, password: string) {
    const body = {
      username: username,
      email: email,
      password: password,
    };

    const response = await this.client.post('/auth/signup', body);
    return response;
  }

  async createMemo(title: string, content: string) {
    const headers = {
      "Auth-Token": this.getAuthToken(),
    };

    const body = {
      title: title,
      content: content,
    };

    const response = await this.client.post('/memos', body, { headers });
    return response;
  }

  async getMemos() {
    const headers = {
      "Auth-Token": this.getAuthToken(),
    };

    const response = await this.client.get('/memos', { headers });
    return response;
  }

  async updateMemo(id: string, title: string, content: string) {
    const headers = {
      "Auth-Token": this.getAuthToken(),
    };

    const body = {
      id: parseInt(id),
      title: title,
      content: content,
    };

    const response = await this.client.put('/memos', body, { headers });
    return response;
  }

  async deleteMemo(id: string) {
    const headers = {
      "Auth-Token": this.getAuthToken(),
    };

    const body = {
      id: parseInt(id),
    };

    const response = await this.client.delete('/memos', { headers, data: body });
    return response;
  }
}

const goServer = new GoServer();
export default goServer;
