type ReplacerFunc = (key: string, value: any) => any;

export const bigintReplacer: ReplacerFunc = (key, value) => {
	return typeof value === "bigint" ? value.toString() : value;
};
