import axios from 'axios';

const USER_API_BASE_URL = 'http://localhost:3001/users';

class UserService {
     async getUsers() {

        try {
            const response =  await axios.get(USER_API_BASE_URL);
            return response.data;
        } catch (error) {
            console.log(error);
            throw error;
        }
    }
     async createUser(user) {
        try {
            const response =  await axios.post(USER_API_BASE_URL, user);
            return response.data;
        } catch (error) {
            console.log(error);
            throw error;
        }
    }
    async getUserById(userId) {
        try {
            const response = await axios.get(USER_API_BASE_URL + '/' + userId);
            return response.data;
        }
        catch (error) {
            console.log(error);
            throw error;
        }
    }
    async updateUser(user, userId) {
        try {
            const response = await axios.put(USER_API_BASE_URL + '/' + userId, user);
                return response.data;
        }
        catch (error) {
            console.log(error);
            throw error;
        }
    }
    async deleteUser(userId) {
        try {
            const response = await axios.delete(USER_API_BASE_URL + '/' + userId);
            return response.data;
        } catch (error) {
            console.log(error);
            throw error;
        }
    }
}

export default new UserService();
