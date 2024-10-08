import { LocksInterface } from "../../interfaces/ILock";
import { ReservesInterface } from "../../interfaces/IReserve";
import { ReserveDetailsInterface } from "../../interfaces/IReserveDetails";
import { UsersInterface } from "../../interfaces/IUser";
import { PaymentInterface } from "../../interfaces/IPayment"
import { SignInInterface } from "../../interfaces/SignIn";


import axios from "axios";

const apiUrl = "http://localhost:8000";

const Authorization = localStorage.getItem("token");

const Bearer = localStorage.getItem("token_type");


const requestOptions = {

  headers: {

    "Content-Type": "application/json",

    Authorization: `${Bearer} ${Authorization}`,

  },

};


async function SignIn(data: SignInInterface) {

  return await axios

    .post(`${apiUrl}/signin`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetPaymentByShopId(id: number) {

  return await axios

    .get(`${apiUrl}/payments/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetLocks() {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  let res = await fetch(`${apiUrl}/locks`, requestOptions)
    .then((res) => {
      if (res.status == 200) {
        return res.json();
      } else {
        return false;
      }
    });

  return res;
}


async function GetUsersById(id: string) {

  return await axios

    .get(`${apiUrl}/user/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetReservesByShopId(id: number) {

  return await axios

    .get(`${apiUrl}/reserves/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetReservesDetailsByReserveId(id: number) {

  return await axios

    .get(`${apiUrl}/reservesdetails/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetShopByUserId(id: string) {

  return await axios

    .get(`${apiUrl}/shopbyuser/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}



async function UpdateUsersById(id: string, data: UsersInterface) {

  return await axios

    .put(`${apiUrl}/user/${id}`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


async function CreateUser(data: UsersInterface) {

  return await axios

    .post(`${apiUrl}/signup`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function CreateReserveDetails(data: ReserveDetailsInterface) {

  return await axios

    .post(`${apiUrl}/createdetails`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function CreateReserve(data: ReservesInterface) {
  return await axios
    .post(`${apiUrl}/createreserve`, data, requestOptions)

    .then((res) => res.data)  // คืนเฉพาะ res.data ซึ่งเป็นข้อมูลที่เราสนใจ

    .catch((e) => {

      return e.response;
    });
}

async function UpdateLocksById(id: string, data: LocksInterface) {

  return await axios

    .put(`${apiUrl}/updatelock/${id}`,data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function ResetLocks() {

  return await axios

    .put(`${apiUrl}/resetlocks`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function CreatePayment(data: PaymentInterface) {

  return await axios

    .post(`${apiUrl}/payments`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

export {

  SignIn,

  GetUsersById,

  UpdateUsersById,

  CreateUser,

  GetShopByUserId,

  CreateReserve,

  CreateReserveDetails,

  GetLocks,

  GetReservesByShopId,

  GetReservesDetailsByReserveId,

  UpdateLocksById,

  ResetLocks,

  GetPaymentByShopId,

  CreatePayment


};