import { FormEvent } from "react";

export const noop = (): void => {};

export const preventDefault = (e: FormEvent): void => e.preventDefault();

export const stopEvent = (e: FormEvent): void => e.stopPropagation();
