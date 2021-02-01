import React from 'react';
import { useTranslation } from 'react-i18next';
import { useParams } from 'react-router-dom';
import { DateTime } from 'luxon';

import Footer from 'components/Footer';
import Header from 'components/Header';
import MainContent from 'components/MainContent';
import PageWrapper from 'components/PageWrapper';
import { NavLink, SecondaryNav } from 'components/shared/SecondaryNav';

import './index.scss';

const AccessibilityRequestDetailPage = () => {
  const { t } = useTranslation('accessibility');
  const { accessibilityRequestId } = useParams<{
    accessibilityRequestId: string;
  }>();
  return (
    <PageWrapper className="accessibility-request">
      <Header />
      <MainContent className="margin-bottom-5">
        <SecondaryNav>
          <NavLink to={`/508/request/${accessibilityRequestId}`}>
            {t('tabs.accessibilityRequests')}
          </NavLink>
        </SecondaryNav>
        <div className="grid-container">
          <h1>Medicare Office of Change Initiative 1.3</h1>
          <div className="grid-row grid-gap-lg">
            <div className="grid-col-8">yo</div>
            <div className="grid-col-4">
              <div className="accessibility-request__side-nav">
                <div className="accessibility-request__other-details">
                  <h3>{t('requestDetails.other')}</h3>
                  <dl>
                    <dt className="margin-bottom-1">
                      {t('intake:fields.submissionDate')}
                    </dt>
                    <dd className="margin-0 margin-bottom-2">
                      {DateTime.fromISO(
                        new Date().toISOString()
                      ).toLocaleString(DateTime.DATE_FULL)}
                    </dd>
                    <dt className="margin-bottom-1">
                      {t('intake:fields.businessOwner')}
                    </dt>
                    <dd className="margin-0 margin-bottom-2">
                      Shane Clark, OIT
                    </dd>
                    <dt className="margin-bottom-1">
                      {t('intake:lifecycleId')}
                    </dt>
                    <dd className="margin-0 margin-bottom-2">X200943</dd>
                  </dl>
                </div>
                <button
                  type="button"
                  className="accessibility-request__remove-request"
                >
                  {t('requestDetails.remove')}
                </button>
              </div>
            </div>
          </div>
        </div>
      </MainContent>
      <Footer />
    </PageWrapper>
  );
};

export default AccessibilityRequestDetailPage;
