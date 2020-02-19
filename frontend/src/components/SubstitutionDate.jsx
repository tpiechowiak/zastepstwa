import React, { useState, useEffect } from 'react';
import axios from 'axios';
import styled from 'styled-components';

import ErrorMessage from "./ErrorMessage";

const SubstitutionDateWrapper = styled.div`
  position: fixed; 
  bottom: 0;
  right: 0;
`;

const SubstitutionDate = () => {
  const [date, setDate] = useState("");
  const [error, setError] = useState(null);

  useEffect(() => {
    axios
      .get(`/api/date`)
      .then(({ data }) => setDate(data.date))
      .catch((err) => setError(String(err)));
  }, []);

  return (
    <>
      <ErrorMessage error={error} />
      <SubstitutionDateWrapper>
        Zastępstwa z dnia: 
        {' '}
        {date}
      </SubstitutionDateWrapper>
    </>
  );
};

export default SubstitutionDate;
