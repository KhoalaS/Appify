export type Result<T, E = Error> =
  | { ok: true; value: T }
  | { ok: false; error: E }

export function newResultOk<T, E = Error>(value: T): Result<T, E> {
  return {
    ok: true,
    value,
  }
}

export function newResultError<T, E = Error>(error: E): Result<T, E> {
  return {
    ok: false,
    error,
  }
}
