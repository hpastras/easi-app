import axios from 'axios';
import { takeLatest, call, put } from 'redux-saga/effects';
import { SystemIntakeForm } from 'types/systemIntake';
import { prepareSystemIntakeForApi } from 'data/systemIntake';
import { fetchSystemIntake, saveSystemIntake } from 'types/routines';
import { Action } from 'redux-actions';
import { updateLastActiveAt } from 'reducers/authReducer';

function putSystemIntakeRequest(formData: SystemIntakeForm) {
  // Make API save request
  const data = prepareSystemIntakeForApi(formData);
  return axios.put(`${process.env.REACT_APP_API_ADDRESS}/system_intake`, data);
}

function* putSystemIntake(action: Action<any>) {
  try {
    yield put(saveSystemIntake.request());
    const response = yield call(putSystemIntakeRequest, action.payload);
    yield put(saveSystemIntake.success(response.data));
    yield put(updateLastActiveAt);
  } catch (error) {
    yield put(saveSystemIntake.failure(error.message));
  } finally {
    yield put(saveSystemIntake.fulfill());
  }
}

function getSystemIntakeRequest(id: string) {
  return axios.get(`${process.env.REACT_APP_API_ADDRESS}/system_intake/${id}`);
}

function* getSystemIntake(action: Action<any>) {
  try {
    yield put(fetchSystemIntake.request());
    const response = yield call(getSystemIntakeRequest, action.payload);
    yield put(fetchSystemIntake.success(response.data));
    yield put(updateLastActiveAt);
  } catch (error) {
    yield put(fetchSystemIntake.failure(error.message));
  } finally {
    yield put(fetchSystemIntake.fulfill());
  }
}

export default function* systemIntakeSaga() {
  yield takeLatest(fetchSystemIntake.TRIGGER, getSystemIntake);
  yield takeLatest(saveSystemIntake.TRIGGER, putSystemIntake);
}
