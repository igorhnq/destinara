const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

interface ApiResponse<T> {
  data?: T;
  error?: string;
  message?: string;
}

class ApiService {
  private baseURL: string;
  private token: string | null;

  constructor() {
    this.baseURL = API_BASE_URL;
    this.token = localStorage.getItem('token');
  }

  private getHeaders(): HeadersInit {
    const headers: HeadersInit = {
      'Content-Type': 'application/json',
    };

    if (this.token) {
      headers['Authorization'] = `Bearer ${this.token}`;
    }

    return headers;
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<ApiResponse<T>> {
    try {
      const url = `${this.baseURL}${endpoint}`;
      const response = await fetch(url, {
        ...options,
        headers: this.getHeaders(),
      });

      const data = await response.json();

      if (!response.ok) {
        return {
          error: data.error || 'Erro na requisição',
        };
      }

      return { data };
    } catch (error) {
      return {
        error: 'Erro de conexão com o servidor',
      };
    }
  }

  // Autenticação
  async register(userData: {
    name: string;
    email: string;
    password: string;
  }) {
    return this.request('/auth/register', {
      method: 'POST',
      body: JSON.stringify(userData),
    });
  }

  async login(credentials: { email: string; password: string }) {
    const response = await this.request('/auth/login', {
      method: 'POST',
      body: JSON.stringify(credentials),
    });

    if (response.data) {
      this.token = response.data.token;
      localStorage.setItem('token', response.data.token);
      localStorage.setItem('user', JSON.stringify(response.data.user));
    }

    return response;
  }

  logout() {
    this.token = null;
    localStorage.removeItem('token');
    localStorage.removeItem('user');
  }

  // Pacotes de viagem
  async getPackages(filters?: {
    available?: boolean;
    destination?: string;
  }) {
    const params = new URLSearchParams();
    if (filters?.available !== undefined) {
      params.append('available', filters.available.toString());
    }
    if (filters?.destination) {
      params.append('destination', filters.destination);
    }

    const endpoint = `/packages${params.toString() ? `?${params.toString()}` : ''}`;
    return this.request(endpoint);
  }

  async getPackageById(id: number) {
    return this.request(`/packages/${id}`);
  }

  // Reservas
  async getUserBookings() {
    return this.request('/bookings');
  }

  async getBookingById(id: number) {
    return this.request(`/bookings/${id}`);
  }

  async createBooking(bookingData: {
    travel_package_id: number;
    start_date: string;
    end_date: string;
  }) {
    return this.request('/bookings', {
      method: 'POST',
      body: JSON.stringify(bookingData),
    });
  }

  async updateBooking(
    id: number,
    bookingData: {
      start_date?: string;
      end_date?: string;
      status?: string;
    }
  ) {
    return this.request(`/bookings/${id}`, {
      method: 'PUT',
      body: JSON.stringify(bookingData),
    });
  }

  async cancelBooking(id: number) {
    return this.request(`/bookings/${id}`, {
      method: 'DELETE',
    });
  }

  // Utilitários
  isAuthenticated(): boolean {
    return !!this.token;
  }

  getCurrentUser() {
    const user = localStorage.getItem('user');
    return user ? JSON.parse(user) : null;
  }
}

export const apiService = new ApiService();
export default apiService; 