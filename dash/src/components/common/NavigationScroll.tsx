import { ReactNode, useEffect } from "react";
import { useLocation } from "react-router-dom";

interface Props {
  children: ReactNode;
}

const NavigationScroll = ({ children }: Props): JSX.Element => {
  const { pathname } = useLocation();
  useEffect(() => {
    window.scrollTo({
      top: 0,
      left: 0,
      behavior: "smooth",
    });
  }, [pathname]);
  return <div>{children || null}</div>;
};

export default NavigationScroll;
