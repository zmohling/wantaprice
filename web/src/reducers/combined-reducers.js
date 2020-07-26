import { userConstants } from '../constants/user-constants';

import authReducer from './authentification';
import registerReducer from './registration';
import { combineReducers } from 'redux';

const allReducers = combineReducers({
    auth: authReducer,
    registration: registerReducer
});

const rootReducer = (state, action) => {
    if (action.type = userConstants.DESTROY_SESSION) {
        state = undefined;
    }
    return allReducers(state, action);
};

export default rootReducer;