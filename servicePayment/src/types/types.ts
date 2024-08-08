export type Transaction = {
  sender: string;
  reciver: string;
  date: string;
  status: Status;
  process_type: ProcessType | string;
  payment_type: string;
  user_id: string;
  amount: number;
  currency: string;
  tenor: number;
};

export type ProcessType = "withdraw" | "transfer";

export type Status = "pending" | "completed" | "failed";
