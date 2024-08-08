import { prisma } from "../clients/prisma/prisma";
import { Transaction } from "../types/types";
import { UpdateAccount } from "./account";
import { CreateLoan, Loan } from "./loan";

export function ProcessTransaction(transaction: Transaction): void {
  try {
    switch (transaction.process_type) {
      case "withdraw":
        Withdraw(transaction)
          .then((result) => {
            console.log("Transaction Success", result);
          })
          .catch((err) => {
            console.log("Transaction Failed", err);
          });
        break;
      case "transfer":
        Transfer(transaction)
          .then((result) => {
            console.log("Transaction Success");
          })
          .catch((err) => {
            console.log("Transaction Failed");
          });
        break;
      case "loan":
        const loan: Loan = {
          user_id: transaction.user_id,
          amount: transaction.amount,
          currency: transaction.currency,
          tenor: transaction.tenor,
        };
        CreateLoan(loan)
          .then((result) => {
            console.log("Loan Success");
          })
          .catch((err) => {
            console.log("Loan Failed");
          });
        break;
      default:
        break;
    }
  } catch (error) {
    console.error("Error in ProcessTransaction:", error);
  }
}

export function GetTransaction(): void {}

async function Withdraw(transaction: Transaction): Promise<void> {
  return new Promise((resolve, reject) => {
    const _ = prisma.tr_account
      .findFirst({
        where: {
          user_id: transaction.user_id,
          type: "default",
        },
      })
      .then((data) => {
        if (!data) {
          console.log("Account not found");
          return;
        }

        if (Number(data.balance) < transaction.amount) {
          console.log("Insufficient Balance");
          return;
        }
        UpdateAccount({
          account_id: data.account_id,
          balance: Number(data.balance) - transaction.amount,
        }).catch((err) => {
          console.log(err);
          console.log("Account not found");
        });
      });

    setTimeout(() => {
      resolve();
    }, 3000);
  });
}

async function Transfer(transaction: Transaction): Promise<void> {
  try {
    const reciever = await prisma.tr_account.findFirst({
      where: {
        user_id: transaction.reciver,
        type: "default",
      },
    });

    const sender = await prisma.tr_account.findFirst({
      where: {
        user_id: transaction.sender,
        type: "default",
      },
    });

    if (!reciever || !sender) {
      console.log("Account not found");
      return;
    }

    await SaveTransaction(transaction);

    await UpdateAccount({
      account_id: reciever.account_id,
      balance: Number(reciever.balance) + transaction.amount,
    });

    await UpdateAccount({
      account_id: sender.account_id,
      balance: Number(sender.balance) - transaction.amount,
    });
  } catch (error) {
    console.error("Error in Transfer:", error);
  }
}

async function SaveTransaction(transaction: Transaction): Promise<void> {
  try {
    await prisma.tr_transaction.create({
      data: {
        sender: transaction.sender,
        receiver: transaction.reciver,
        amount: transaction.amount,
        currency: transaction.currency,
        type: transaction.process_type,
        user_id: transaction.sender,
        status: "success",
      },
    });
  } catch (error) {
    console.log("Error", error);
  }
}
