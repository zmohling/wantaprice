import { userConstants } from '../constants/user-constants';

const registerReducer = (state = {}, action) => {
    switch (action.type) {
        case userConstants.REGISTER_REQUEST:
            return {
                registering: true
            };
        case userConstants.REGISTER_SUCCESS:
            return {};
        case userConstants.REGISTER_FAILURE:
            return {};
        default:
            return state;
    }
};

export default registerReducer;