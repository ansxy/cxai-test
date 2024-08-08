import { prisma } from "../clients/prisma/prisma";

type GetAccount = {
  user_id: string;
  type?: string;
};

export async function findAccount(account: GetAccount) {
  return new Promise((resolve, reject) => {
    const data = prisma.tr_account.findFirst({
      where: {
        user_id: account.user_id,
        type: account.type || "",
      },
    });

    if (!data) {
      reject("Account not found");
    }

    resolve(data);
  });
}

type ReqUpdateAccount = {
  account_id: string;
  balance: number;
};

export async function UpdateAccount(req: ReqUpdateAccount) {
  return new Promise((resolve, reject) => {
    const data = prisma.tr_account.update({
      where: {
        account_id: req.account_id,
      },
      data: {
        balance: req.balance,
      },
    });

    if (!data) {
      reject("Account not found");
    }

    resolve(data);
  });
}

module.exports = {
  findAccount,
  UpdateAccount,
};
