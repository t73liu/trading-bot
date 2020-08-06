// eslint-disable-next-line import/prefer-default-export
export const getJSON = async (url) => {
  try {
    const response = await fetch(url, {
      headers: { "Content-Type": "application/json" },
    });
    if (response.ok) {
      return await response.json();
    }
    const contentType = response.headers.get("Content-Type");
    if (contentType?.includes("application/json")) {
      const json = await response.json();
      return new Error(json);
    }
    const text = await response.text();
    return new Error(text);
  } catch (e) {
    return e;
  }
};
