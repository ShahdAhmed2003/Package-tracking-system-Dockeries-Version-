export const decodeToken = (token) => {
    const base64URL = token.split('.')[1];
    const base64 = base64URL.replace('-', '+').replace('_', '/');
    const decodedToken = JSON.parse(atob(base64));
    return decodedToken;
};
