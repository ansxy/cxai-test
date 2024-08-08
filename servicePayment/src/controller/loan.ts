import { prisma } from "../clients/prisma/prisma";

export type Loan = {
  user_id: string;
  amount: number;
  currency: string;
  tenor: number;
};
export async function CreateLoan(loan: Loan): Promise<void> {
  try {
    const createdLoan = await prisma.tr_loan.create({
      data: {
        user_id: loan.user_id,
        amount: loan.amount,
        currency: loan.currency,
        tenor: loan.tenor,
        interest_rate: 0.05,
      },
    });

    // Generate loan logs
    const loanLogs = [];
    for (let i = 0; i < loan.tenor; i++) {
      loanLogs.push({
        loan_id: createdLoan.loan_id,
        total_amount: loan.amount / loan.tenor,
        is_paid: false,
        duedate: new Date(new Date().setMonth(new Date().getMonth() + i + 1)),
      });
    }

    await prisma.tr_loan_log.createMany({
      data: loanLogs,
    });

    const userid = await prisma.tr_account.findFirst({
      where: {
        user_id: loan.user_id,
        type: "default",
      },
    });

    if (!userid) {
      console.log("Account not found");
      return;
    }

    await prisma.tr_account.update({
      where: {
        account_id: userid.account_id,
      },
      data: {
        balance: Number(userid.balance) + loan.amount,
      },
    });
  } catch (error) {
    console.error("Error creating loan and loan logs:", error);
    throw error;
  }
}

module.exports = { CreateLoan };
