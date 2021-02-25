import react from 'react';

const ExportCard = () => {
    return (
        <div className="card">
            <i className="fa fa-download fa-2x text-red" aria-hidden="true"></i>
            <div className="card_inner">
                <p className="text-primary-p">Export contacts...</p>
            </div>
        </div>
    );
}

export default ExportCard;