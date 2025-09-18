import toast from "react-hot-toast";

export const notify = {
  success: (msg, delay = 1000) => setTimeout(() => toast.success(msg), delay),

  error: (msg, delay = 1000) => setTimeout(() => toast.error(msg), delay),

  info: (msg, delay = 1000) => setTimeout(() => toast(msg), delay),
};


